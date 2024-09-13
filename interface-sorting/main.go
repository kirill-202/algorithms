package main

import (
	"fmt"
)

/*
Description:
Create a Go program that sorts a collection of different types of numbers using interfaces. The program should support sorting integers and floating-point numbers,
 demonstrating the use of interfaces to handle different number types.
 */

type Sortable interface{
	Less(i, i2 int) bool
	Swap(i, i2 int)
	Sort()

}

type SortableGen interface{
	~int | ~float64

}

type MyIntSlice []int

type MyFloatSlice []float64

func (mf MyFloatSlice) Less(i, i2 int) bool {
	return mf[i] < mf[i2]
}

func (mf MyFloatSlice) Swap(i, i2 int) {
	mf[i], mf[i2] = mf[i2], mf[i]
}

func (mf MyFloatSlice) Sort() {
	length := len(mf)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if !mf.Less(j, j+1) {
				mf.Swap(j, j+1)
			}
		}
	}
}



func (mi MyIntSlice) Less(i, i2 int) bool {
	return mi[i] < mi[i2]
}

func (mi MyIntSlice) Swap(i1, i2 int) {
	mi[i1], mi[i2] = mi[i2], mi[i1]
}

func (mi MyIntSlice) Sort() {
	length := len(mi)
    for i := 0; i < length-1; i++ { 
        for j := 0; j < length-i-1; j++ {
            if !mi.Less(j, j+1) {
                mi.Swap(j, j+1)
			}
		}
	}
}

func Sort[T SortableGen](slice []T) {
	length := len(slice)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if slice[j] < slice[j+1] { 
				slice[j], slice[j+1] = slice[j+1], slice[j] 
			}
		}
	}
}




func main() {
	var testInt MyIntSlice
	for i:=10; i > 0; i-- {
		testInt = append(testInt, i)
	}
	fmt.Println(testInt)
	testInt.Sort()
	fmt.Println(testInt)
	Sort(testInt)
	fmt.Println(testInt)

}