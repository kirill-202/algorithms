package main

import (
	"fmt"
)

/*
Task Details:
You need to implement the Merge Sort algorithm, which recursively splits the array into smaller subarrays, sorts them, and then merges the sorted subarrays back together.
The program should handle arrays of varying lengths, including empty arrays.
You should account for edge cases such as arrays with all identical elements or arrays that are already sorted.
Make sure to include helper functions for merging subarrays.
Input:
An unsorted array of integers, e.g., [34, 7, 23, 32, 5, 62].
*/

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2

	left := MergeSort(array[:mid])
	right := MergeSort(array[mid:])
	fmt.Printf("Left %v, Right %v\n", left, right)
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int

	indexL, indexR := 0, 0

	for indexL < len(left) && indexR < len(right) {
		if left[indexL] < right[indexR] {
			result = append(result, left[indexL])
			indexL++

		} else {
			result = append(result, right[indexR])
			indexR++
		}
	}
	result = append(result, left[indexL:]...)
	result = append(result, right[indexR:]...)
	return result

}

func main() {

	testInput := []int{34, 7, 23, 32, 5, 62, 4, 86, 7}
	result := MergeSort(testInput)
	fmt.Println("Start...", result)
}
