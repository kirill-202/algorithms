package main


import (
	"fmt"
)


func BinarySearch(arr []int, val int) int {
    if len(arr) == 0 {
        return -1
    }

	pivot := len(arr)/2

    if arr[pivot] == val {
        return pivot
	} else if val < arr[pivot] {
		return BinarySearch(arr[:pivot], val)
	} else {
		res := BinarySearch(arr[pivot+1:], val)
		if res == -1 {
			return -1
		}
		return res + pivot + 1
			}

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