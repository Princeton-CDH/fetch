// CLI to scrape a website recursively, collecting information on links
// and how many times they are internally referenced, as well as source, status
// code, etc.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"

	"github.com/gocolly/colly"

)



func main() {

	// configure a maxDpeth and max Routines set of flags
	maxDepth := flag.Int("max-depth", 0, "maximum scraping depth from initial")
	maxRoutines := flag.Int("max-routines", 0, "maximum number of go Routines")

	// parse flags
	flag.Parse()

	// look for arguments and note the problem, bail out on missing url to scrape
	args := flag.Args()
	if len(args) > 1 {
		fmt.Printf("Ignoring extra args %v\n", args[1:])
	} else if len(args) == 0 {
		fmt.Printf("Provide a url to scrape.\n")
		os.Exit(1)
	}
	site := args[0]
	c := ConfigureColly(site, *maxRoutines, *maxDepth)

	// map to hold the Urls collected by the various goroutines
	urls := Urls{m: make(map[string]*Url)}


	c.OnResponse(func(r *colly.Response) {
		// get information for a URL that appears in the response
		urlPath := r.Request.URL.String()
		// lock the mutex to safely modify the map and defer an unlock
		urls.Lock()
		defer urls.Unlock()

		// update the URL if it exists already (and most will having been added in
		// onHTML)
		if _, ok := urls.m[urlPath]; ok {
			url := urls.m[urlPath]
			url.Url = urlPath
			url.StatusCode = r.StatusCode
			url.LastModified = r.Headers.Get("Last-Modified")
			url.Size = len(r.Body)
			cl, err := strconv.Atoi(r.Headers.Get("Content-Length"))
			if err == nil {
				url.ContentLength = cl
			}
			url.Timestamp = time.Now().UTC().Format(time.RFC1123)
		} else {
			// otherwise initialize it for the first time
			url := &Url{
				Url: urlPath,
				SourceUrl: "",
				LinkCount: 1,
				StatusCode: r.StatusCode,
				LastModified: r.Headers.Get("Last-Modified"),
				Size: len(r.Body),
				Timestamp: time.Now().UTC().Format(time.RFC1123),
			}
			cl, err := strconv.Atoi(r.Headers.Get("Content-Length"))
			if err != nil {
				url.ContentLength = cl
			}
			urls.m[urlPath] = url
		}
	})


	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		sourceUrl := e.Request.URL.String()
		urls.Lock()

		if strings.HasPrefix(e.Attr("href"), "mailto") {
			urls.Unlock()
			return
		}

		link := e.Request.AbsoluteURL(e.Attr("href"))

		inDomain := false
		for _, domain := range c.AllowedDomains {
			if strings.Contains(link, domain) {
				inDomain = true
			}
		}

		if !inDomain {
			urls.Unlock()
			return
		}

		if _, ok := urls.m[link]; ok {
			urls.m[link].LinkCount += 1
			urls.Unlock()
		} else {
			url := &Url{
				Url: link,
				SourceUrl: sourceUrl,
				LinkCount: 1,
			}
			urls.m[link] = url
			urls.Unlock()
			e.Request.Visit(link)
		}
	})


	c.Visit(site)
	c.Wait()
	printResults(&urls)

}
