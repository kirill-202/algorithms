package main


import (
	"fmt"
)


func YieldPivotIndex(arr []int) (pIndex int) {
	pIndex = 0

	if len(arr) < 2 {
        return -1
    }
	for i:=1; i < len(arr)-1; i++ {
		if arr[i] >= arr[pIndex] {
			pIndex = arr[i]
		} else if arr[i+1] { 

		}
	}
}

func main() {
	testAr := []int{3,7,4,5,8,95,17,18,20}
	fmt.Println(testAr)
}