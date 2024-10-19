package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    var active chan int

    go func() {
        time.Sleep(time.Second * 2)
        ch1 <- 100 // Send to ch1 after 2 seconds
		ch1 <- 300
    }()

    go func() {
        time.Sleep(time.Second * 4)
        ch2 <- 200 // Send to ch2 after 4 seconds
    }()

    // Disable ch1 by setting it to nil
	active = ch1

	for i := 0; i < 3; i++ {
		select {
		case val := <-active:
			fmt.Println("Received from ch1:", val)
			// After reading from active channel, disable it
			fmt.Println("Switch to ch2")
			active = ch2
		case val := <-active:
			fmt.Println("Received from ch2:", val)
			// Once we receive from ch2, enable ch1 again
			fmt.Println("Switch to ch1")
			active = ch1
		}
	}

	fmt.Println("Done")
}