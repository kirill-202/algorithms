package main

import (
	"fmt"
)

const TIMES = 3

func MergeSort(array []int) []int {
	subArrays := split(array, TIMES)
	var finArray []int
	for _, value := range subArrays {
		sortChunk(value, len(value))
		finArray = append(finArray, value...)
	}
	sortChunk(finArray, len(finArray))
	return finArray

}

func split(array []int, times int) (chunks [][]int) {
	initIndex := len(array) / times
	for i := 0; i < len(array); i += initIndex {
		end := i + initIndex
		if end > len(array) {
			end = len(array)
		}
		chunks = append(chunks, array[i:end])
	}
	return
}

func sortChunk(chunk []int, n int) {
	if n == 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		if chunk[i] > chunk[i+1] {
			chunk[i], chunk[i+1] = chunk[i+1], chunk[i]
		}
	}
	sortChunk(chunk, n-1)

}

func main() {

	testInput := []int{34, 7, 23, 32, 5, 62, 4, 86, 7}
	result := MergeSort(testInput)
	fmt.Println("Start...", result)
}
