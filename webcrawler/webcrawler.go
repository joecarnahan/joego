package webcrawler

import "fmt"

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Receives URL strings on the input channel, sends unique URL strings
// on the output channel. In other words, each URL might be given to
// "input" multiple times, but it will only be sent to "output" once.
//
// This function runs until as many unregister messages have been received
// as register messages.
func ManageUrls(input chan string, output chan string,
	register chan bool, unregister chan bool) {
	registerCount := 0
	urls := make(map[string]bool)
	for {
		select {
		case url := <-input:
			if !urls[url] {
				urls[url] = true
				output <- url
			}
		case <-register:
			registerCount++
		case <-unregister:
			registerCount--
			if registerCount <= 0 {
				return
			}
		}
	}
}

func WriteUrl(output chan string, url string) {
	output <- url
}

func ParallelCrawl(url string, url_output chan string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
}

// TODO(jcarnahan): Figure out how to safely create and register crawlers

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	register := make(chan bool)
	unregister := make(chan bool)
	uniqueUrls := make(chan string)
	allUrls := make(chan string)
	go WriteUrl(allUrls, url)
	go ManageUrls(allUrls, uniqueUrls, register, unregister)
	for uniqueUrl := range uniqueUrls {
		go ParallelCrawl(uniqueUrl, allUrls, depth-1, fetcher)
	}

	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	return
}
