package main

import (
  "testing"
)

func TestPrintResults(t *testing.T) {
  urls := Urls{
      m: map[string]*Url{
        "http://foobar.org": &Url{
          "http://foobar.org",
          "http://foobar.org",
          1,
          200,
          "Today",
          35,
          35,
          "Today at 3:00",
        },
        "http://foobar.org/boo/": &Url{
          "http://foobar.org/boo/",
          "http://foobar.org",
          1,
          302,
          "Today",
          35,
          35,
          "Today at 3:00",
        },
      },
    }
    printResults(&urls)
    // Output:
    // url,source_url,link_count,status_code,lastModified,contentLength,size,timestamp
    // http://foobar.org,1,200,Today,35,35,Today at 3:00
    // http://foobar.org/boo/,1,302,Today,35,35,Today at 3:00
}
