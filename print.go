package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func printResults(urls *Urls) {
	results := []*Url{}
	for _, v := range urls.m {
		results = append(results, v)
	}
	
	w := csv.NewWriter(os.Stdout)
	w.Write([]string{
			"url", "source_url", "link_count",
			"status_code", "lastModified", "contentLength", "size", "timestamp",
		})
	w.Flush()


	for _, result := range results {
		w.Write([]string{
			result.url,
			strconv.Itoa(result.linkCount),
			strconv.Itoa(result.statusCode),
			result.lastModified,
			strconv.Itoa(result.contentLength),
			strconv.Itoa(result.size),
			result.timestamp,
		})
		w.Flush()
	}

}
