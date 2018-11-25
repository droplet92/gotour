package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var isFetched map[string]bool = make(map[string]bool)

func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if isFetched[url] {
		return
	}
	isFetched[url] = true

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	var wait sync.WaitGroup
	wait.Add(len(urls))

	for _, u := range urls {
		go func(url string) {
			defer wait.Done()
			Crawl(url, depth-1, fetcher)
		}(u)
	}
	wait.Wait()
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

/* belows are implementations of fake fetcher and results */
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
