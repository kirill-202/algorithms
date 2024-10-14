package main


import (
	"fmt"
)


func BinarySearch(arr []int, val int) int {

	low, high := 0, len(arr)-1


	for low <= high {

		mid := (low+high)/2

		if val == arr[mid] {
			return mid
		}
		if val < arr[mid] {
			high = mid-1 
		} else {
			low = mid+1
		}
	}
	return -1

}


func main() {

	sortedArr := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		8,
		10,
		100,
	}

	fmt.Println("Programm has started...")
	fmt.Printf("My sorted array %v\n", sortedArr)

	testValues := []int{3,6,7,10,12,100}
	for _, v := range testValues{
		fmt.Printf("The search result for value %d... Index in array:  %d\n", v, BinarySearch(sortedArr, v))
	}
}