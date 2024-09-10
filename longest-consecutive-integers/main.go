package main

/*
Given an unsorted array of integers,
 find the length of the longest consecutive elements sequence.
 Input: [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. 
Therefore its length is 4.
*/

import (
	"fmt"
	"slices"
)


func isOneIncremented(current, next int) bool {
	return  current+1 == next
}

func formSequances(slc []int) (sslc [][]int) {
	var tempSlice []int
	for i, v := range slc {
		if len(tempSlice) == 0 {
			tempSlice = append(tempSlice, v)
			continue
		}
		if isOneIncremented(slc[i-1], v) {
			tempSlice = append(tempSlice, v)
		} else {
			sslc = append(sslc, tempSlice)
			tempSlice = []int{v}
		}
	}
	sslc = append(sslc, tempSlice)
	return
}

func FindConsLength(slc []int) (slicedSequence []int, length int) {
	slices.Sort(slc)
	trunc := truncDuplicates(slc)
	slicesConsec := formSequances(trunc)
	for _, slice := range slicesConsec {

		if len(slice) > len(slicedSequence) {
			slicedSequence = slice
			length = len(slicedSequence)
		}

	}
	return
}

func truncDuplicates(slc []int) (trunckedSlice []int) {
	
	for _, i := range slc {
		if slices.Contains(trunckedSlice, i) {
			continue
		}
		trunckedSlice = append(trunckedSlice, i)
	}
	return trunckedSlice
}

func main() {
	ts := []int{100, 4, 200, 1, 3, 2, 2, 101, 105}
	fmt.Printf("This is the initial slice %v with the length %d.\n", ts, len(ts))
	winner, wlength := FindConsLength(ts)
	fmt.Printf("This is the slice %v with the length %d.\n", winner, wlength)
}
