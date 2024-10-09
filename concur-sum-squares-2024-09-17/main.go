package main


import (
	"fmt"
	"sync"
	"math"
)

var wg sync.WaitGroup
const chunkSize = 5



func RegularLoopSquare(input []int) float64 {

	var sum float64

	for _, num := range input{
		sum += math.Sqrt(float64(num))
	}

	return sum
}


func ConcurrentLoopSquare(input []int) (finSum float64) {

	wg.Add(len(input))
	ch := make(chan float64, len(input))

	for _, elem := range input {
		go rootSquare(elem, ch)
	}
	go func() {
		wg.Wait()
		defer close(ch)
	}()

	for result := range ch {
		finSum += result
	}

	return 
	
}


func ConcurrentLoopSquareTwo(input []int) float64 {
	ch := make(chan float64, len(input)/chunkSize+1)

	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize
		if end > len(input) {
			end = len(input)
		}
		wg.Add(1)
		go rootSquareChunk(input[i:end], ch)
	}


	go func() {
		wg.Wait()
		close(ch)
	}()


	return rootSum(ch)
}

func rootSquareChunk(chunk []int, ch chan float64) {
	defer wg.Done()
	var chunkSum float64
	for _, elem := range chunk {
		chunkSum += math.Sqrt(float64(elem))
	}
	ch <- chunkSum
}

func rootSum(ch <-chan float64) float64 {
	var sum float64
	for result := range ch {
		sum += result
	}
	return sum
}


func rootSquare(elem int, ch chan float64) {
	defer wg.Done()

	ch <- math.Sqrt(float64(elem))
}





func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(RegularLoopSquare(input))
	fmt.Println(ConcurrentLoopSquare(input))
	fmt.Println()


}