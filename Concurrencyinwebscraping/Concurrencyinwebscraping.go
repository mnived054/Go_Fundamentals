package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func scraper(wg *sync.WaitGroup, url string, results chan string) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Failed to fetch %s: %v", url, err)
		return
	}
	body, _ := io.ReadAll(response.Body)
	results <- fmt.Sprintf("Fetched %s\n: %s bytes", url, body)
}

func main() {
	var wg sync.WaitGroup
	sites := []string{
		"https://www.google.com/",
		"https://www.golang.org",
	}
	results := make(chan string, len(sites))
	wg.Add(len(sites))

	for _, site := range sites {
		go scraper(&wg, site, results)
	}

	wg.Wait()
	close(results)
	for result := range results {
		fmt.Println(result)
	}
}
