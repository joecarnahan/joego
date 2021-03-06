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

type page struct {
	url   string
	depth int
}

// Reads pages from allPages and writes out any pages that should be crawled
// based on what pages have been seen so far. Specifically, the same URL will
// only be crawled twice if it is encountered again at a higher depth than it
// was encountered before.
//
// Loops until the allPages channel is closed. Closes the filteredPages channel
// before exiting.
func filterPages(allPages chan page, filteredPages chan page) {
	urlDepths := make(map[string]int)
	for next_page := range allPages {
		if urlDepths[next_page.url] < next_page.depth {
			urlDepths[next_page.url] = next_page.depth
			filteredPages <- next_page
		}
	}
	close(filteredPages)
}

// Prints any strings that are sent to it to the given Writer. Loops until the
// given channel is closed.
func printStrings(toPrint chan string, output io.Writer) {
	for s := range toPrint {
		io.WriteString(output, s)
	}
}

// Fetches the page at the given URL, prints its output to the given text
// output channel, writes all of the URLs on the page to the given page output
// channel (decrementing depth by 1), and then writes to the given boolean
// channel when done.
func parallelCrawl(fetcher Fetcher, to_crawl page, text_output chan string, page_output chan page, done chan bool) {
	body, urls, err := fetcher.Fetch(to_crawl.url)
	if err != nil {
		text_output <- fmt.Sprintf("%v\n", err)
		done <- true
		return
	}
	text_output <- fmt.Sprintf("found: %s %q\n", to_crawl.url, body)
	for _, url := range urls {
		page_output <- page{url, to_crawl.depth - 1}
	}
	done <- true
}

// Sets up a pipeline of pages to crawl seeded with the given URL and crawls
// them in parallel using the given URL fetcher.
func Crawl(url string, depth int, fetcher Fetcher) {
	// Set up threadsafe output.
	output := make(chan string)
	go printStrings(output, os.Stdout)

	// Set up pipeline of pages to crawl.
	allPages := make(chan page)
	filteredPages := make(chan page)
	go filterPages(allPages, filteredPages)
	allPages <- page{url, depth}

	// Count the number of concurrent crawls to know when we are done.
	crawlers := 0
	doneCrawling := make(chan bool)

	for {
		select {
		case next_page := <-filteredPages:
			crawlers++
			go parallelCrawl(fetcher, next_page, output, allPages, doneCrawling)
		case <-doneCrawling:
			crawlers--
			if crawlers == 0 {
				return
			}
		}
	}
}
