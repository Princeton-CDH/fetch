package main

import (
	"sync"
)

type Url struct {
	Url string
	SourceUrl string
	LinkCount int
	StatusCode int
	LastModified string
	Size int
	ContentLength int
	Timestamp string
}

type Urls struct {
	sync.RWMutex
	m map[string]*Url
}
