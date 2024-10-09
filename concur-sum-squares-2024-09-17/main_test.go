package main

import (
	"testing"
)





func BenchmarkRegularLoopSquare(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for i := 0; i < b.N; i++ {
		RegularLoopSquare(input)
	}
}


func BenchmarkConcurrentLoopSquare(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for i := 0; i < b.N; i++ {
		ConcurrentLoopSquare(input)
	}
}

func BenchmarkConcurrentLoopSquareTwo(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for i := 0; i < b.N; i++ {
		ConcurrentLoopSquareTwo(input)
	}
}




func BenchmarkRegularLoopSquare1000(b *testing.B) {
	hugeInput := func(n int) []int {
		s := make([]int, n)
		for i := 0; i < n; i++ {
			s[i] = i
		}
		return s
	}(1000)
	
	
	for i := 0; i < b.N; i++ {
		RegularLoopSquare(hugeInput)
	}
}


func BenchmarkConcurrentLoopSquare1000(b *testing.B) {
	hugeInput := func(n int) []int {
		s := make([]int, n)
		for i := 0; i < n; i++ {
			s[i] = i
		}
		return s
	}(1000)
	
	for i := 0; i < b.N; i++ {
		ConcurrentLoopSquare(hugeInput)
	}
}


func BenchmarkConcurrentLoopSquareTwo1000(b *testing.B) {
	hugeInput := func(n int) []int {
		s := make([]int, n)
		for i := 0; i < n; i++ {
			s[i] = i
		}
		return s
	}(1000)
	
	for i := 0; i < b.N; i++ {
		ConcurrentLoopSquareTwo(hugeInput)
	}
}

