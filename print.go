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
			result.Url,
			strconv.Itoa(result.LinkCount),
			strconv.Itoa(result.StatusCode),
			result.LastModified,
			strconv.Itoa(result.ContentLength),
			strconv.Itoa(result.Size),
			result.Timestamp,
		})
		w.Flush()
	}

}
