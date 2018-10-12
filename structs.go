package main

import (
	"sync"
)

type Url struct {
	url string
	sourceUrl string
	linkCount int
	statusCode int
	lastModified string
	size int
	contentLength int
	timestamp string
}

type Urls struct {
	sync.RWMutex
	m map[string]*Url
}
