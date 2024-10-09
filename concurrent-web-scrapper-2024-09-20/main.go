package main

/*

Task: Concurrent Web Scraper
Problem Description:
You need to implement a concurrent web scraper that fetches content from a list of URLs. The goal is to scrape multiple websites concurrently, process their content, and aggregate the results. You'll need to practice using channels and goroutines effectively.

Requirements:
Input: A list of URLs provided as input.
Concurrency: You need to fetch the content of each URL concurrently using goroutines.
Channels: Use channels to send the content (or some processed result, like the length of the content) from the goroutines back to the main routine.
Worker Pool: Implement a worker pool with a limited number of workers to control concurrency.
Error Handling: Handle errors gracefully when a URL cannot be fetched (e.g., due to a timeout or unreachable server).
Timeouts: Set a timeout for the HTTP requests to ensure that the program doesn’t hang if a website is too slow.
Output: Print the number of characters fetched from each URL or the status of the operation.
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const LIMIT = 5

var activeWorkers int32

type Worker struct {
	Client *http.Client
}

var workerPool = sync.Pool{
	New: func() interface{} {
		return &Worker{
			Client: &http.Client{},
		}
	},
}

func (w *Worker) Get(url string) {
	response, err := w.Client.Get(url)

	if err != nil {
		fmt.Printf("error getting the data via %s, %v\n", url, err)
		return
	}
	defer response.Body.Close()
	fmt.Printf(" Prcessing data from %s... Status code %v. fetching data...\n", url, response.StatusCode)
	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("error reading body, %v\n", err)
	}
	domain := trimUrl(url)
	if err := SaveToFile(body, domain, "lending_pages"); err != nil {
		fmt.Printf("error saving data to file, %e\n", err)
	}
	fmt.Printf("The content size for %s is %d\n", domain, len(body))

}

func SaveToFile(body []byte, filename, subDirectory string) error {
	txtFN := subDirectory + "/" + filename + ".txt"
	err := os.Mkdir(subDirectory, 0770)
	if os.IsNotExist(err) {
		fmt.Println("Creating a directory...")
	}
	file, err := os.Create(txtFN)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(body); err != nil {
		return err
	}
	return nil
}

func trimUrl(url string) string {
	trimmedSu := strings.TrimSuffix(url, ".com")
	fullTrim := strings.TrimPrefix(trimmedSu, "https://")
	return fullTrim
}

func ProcessTask(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	defer atomic.AddInt32(&activeWorkers, -1)
	worker := workerPool.Get().(*Worker)

	defer workerPool.Put(worker)
	worker.Get(url)

}

func main() {
	wg := sync.WaitGroup{}
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://github.com",
		"https://yandex.com",
		"https://pkg.go.dev",
	}
	for _, url := range urls {

		for atomic.LoadInt32(&activeWorkers) >= LIMIT {
			time.Sleep(100 * time.Millisecond)
		}

		wg.Add(1)
		atomic.AddInt32(&activeWorkers, 1)

		go ProcessTask(url, &wg)

		fmt.Printf("Active workers: %d\n", atomic.LoadInt32(&activeWorkers))

	}

	wg.Wait()
}
