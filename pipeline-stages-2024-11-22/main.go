package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const BUFFER = 10

func GenNums(passTo chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	num := rand.Intn(50)
	fmt.Println("Number generated:", num)
	passTo <- num
	time.Sleep(1 * time.Second)
}

func SquareNum(passFrom chan int, passTo chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	num := <-passFrom
	sqrtNum := math.Sqrt(float64(num))
	fmt.Printf("Square root for %d = %.2f\n", num, sqrtNum)
	passTo <- sqrtNum
	time.Sleep(2 * time.Second)
}

func FilterEven(passFrom chan float64, passTo chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	num := <-passFrom
	if int(num)%2 == 0 {
		fmt.Println("Even number:", num)
		passTo <- num
	}
	time.Sleep(3 * time.Second)
}

func main() {
	chan1 := make(chan int, BUFFER)
	chan2 := make(chan float64, BUFFER)
	chan3 := make(chan float64, BUFFER)

	var wg sync.WaitGroup

	for i := 0; i < BUFFER; i++ {
		wg.Add(1)
		go GenNums(chan1, &wg)

		wg.Add(1)
		go SquareNum(chan1, chan2, &wg)

		wg.Add(1)
		go FilterEven(chan2, chan3, &wg)
	}

	wg.Wait()

	close(chan1)
	close(chan2)
	close(chan3)

	var result []float64
	for item := range chan3 {
		result = append(result, item)
	}

	fmt.Println("This is the result of iterations:", result)
}
