package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	runs       = 40
	numWorkers = 3
)

func WaitWriter(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range ch {

		start := time.Now()
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		elapsed := time.Since(start).Seconds()
		fmt.Printf("The process number %d, with time to run for  %f\n", value, elapsed)
	}
}

func main() {
	var wg sync.WaitGroup
	intCh := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go WaitWriter(intCh, &wg)

	}
	for i := 1; i < runs+1; i++ {
		intCh <- i
	}

	defer func() {
		close(intCh)
		wg.Wait()
	}()

}
