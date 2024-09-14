package main


import (
	"fmt"
	"math/rand"
)


/*
Given an integer array nums and an integer k
, return the k-th largest element in the array.

Note: You need to find the k-th largest element in the sorted order,
 not the k-th distinct element.
 */


 type MaxHeap struct {
	slice []int
}


func (h *MaxHeap) Insert(key int) {
	h.slice = append(h.slice, key)
	h.heapifyUp(len(h.slice) - 1)
}

func (h *MaxHeap) Extract() int {
	if len(h.slice) == 0 {
		fmt.Println("No elements in heap")
		return -1
	}

	extracted := h.slice[0]
	h.slice[0] = h.slice[len(h.slice)-1]
	h.slice = h.slice[:len(h.slice)-1]

	h.heapifyDown(0)

	return extracted
	
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.slice[parent(index)] < h.slice[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.slice) - 1
	left, right := leftChild(index), rightChild(index)
	childToCompare := 0

	for left <= lastIndex {
		if left == lastIndex {
			childToCompare = left
		} else if h.slice[left] > h.slice[right] { 
			childToCompare = left
		} else { 
			childToCompare = right
		}

		if h.slice[index] > h.slice[childToCompare] {
			return
		}

		h.swap(index, childToCompare)
		index = childToCompare
		left, right = leftChild(index), rightChild(index)
	}
}

func (h *MaxHeap) FindKthInt(k int) int {

	if k > len(h.slice) { 
		return -1 
	}

	var TempExtract int
	for i:=0; i < k; i++ {
		TempExtract = h.Extract()
	}
	return TempExtract
}


func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.slice[i1], h.slice[i2] = h.slice[i2], h.slice[i1]
}



 func main() {
	var testSlice []int
	for i := 0; i < 10; i++ {
		testSlice = append(testSlice, rand.Intn(1000))
	}

	mh := &MaxHeap{}
	for _, v := range testSlice {
		mh.Insert(v)
	}

	fmt.Println(mh.slice)
	kValue := mh.FindKthInt(12)
	fmt.Println(kValue)
 }