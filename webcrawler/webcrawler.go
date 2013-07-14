package webcrawler

import (
	"fmt"
	"io"
	"os"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type result struct {
	url string
	depth int
}

// Reads results from allResults and writes out any results that should be crawled based
// on what results have been seen so far. Specifically, the same URL will only be crawled
// twice if it is encountered again at a higher depth than it was encountered before.
//
// Loops until the allResults channel is closed. Closes the filteredResults channel before
// exiting.
func filterResults(allResults chan result, filteredResults chan result) {
	urlDepths := make(map[string]int)
	for result := range allResults {
		if (urlDepths[result.url] == 0) || /* RESUME HERE Figure out depth calculation */ (urlDepths[result.url] {
			urlDepths[result.url] = result.depth
			filteredResults <- result
		}
	}
}

// Prints any strings that are sent to it to the given Writer. Loops until the given
// channel is closed.
func printStrings(toPrint chan string, output io.Writer) {
	for s := range toPrint {
		io.WriteString(output, s)
	}
}

func parallelCrawl(url string, depth int, 

func Crawl(url string, depth int, fetcher Fetcher) {
	output := make(chan string)
	go printStrings(output, os.Stdout)
	allUrls := make(chan string)
	filteredUrls := make(chan string)
	go filterUrls(allUrls, filteredUrls)
	// Set up URL filter
	// Set up printer
	// Set up crawler registry
	// Send first URL to URL filter




















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
	filteredUrls := make(chan string)
	allUrls := make(chan string)
	go WriteUrl(allUrls, url)
	go ManageUrls(allUrls, filteredUrls, register, unregister)
	for uniqueUrl := range filteredUrls {
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
