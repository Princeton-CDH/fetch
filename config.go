package main

import (
	"fmt"
	"net/url"
	"os"
	"github.com/gocolly/colly"

)

func ConfigureColly(site string, maxRoutines, maxDepth int) (*colly.Collector) {

	siteUrl, err := url.Parse(site)

	if err != nil {
		fmt.Println("Malformed or unparseable url.")
		os.Exit(1)
	}

	args := []func(*colly.Collector){
		colly.AllowedDomains(siteUrl.Hostname()),
	}

	if maxDepth > 0 {
		args = append(args, colly.MaxDepth(maxDepth))
	}

	if maxRoutines > 0 {
		args = append(args, colly.Async(true))
	}

	c := colly.NewCollector(args...)

	if maxRoutines > 0 {
		c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: maxRoutines})
	}
	return c
}
