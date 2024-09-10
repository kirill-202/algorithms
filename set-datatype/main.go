package main

import (
	"fmt"
)



type Set[T comparable] struct{
	data map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
        data: make(map[T]bool),
    }
}

func (s *Set[T]) Add(value T) {
	s.data[value] = true
}

func (s *Set[T]) Remove(value T) {
	delete(s.data, value)
}

func (s *Set[T]) Contains(value T) bool {
	return s.data[value]
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) Elements()  []T {
	backSlice := make([]T, 0, len(s.data))
	for key := range s.data {
		backSlice = append(backSlice, key)
	}
	return backSlice
}


func main() {
	mySet := NewSet[int]()

	for i:=0; i<10; i++ {
		mySet.Add(i)
	}
	for i:=0; i<10; i++ {
		mySet.Add(i)
	}
	mySet.Remove(3)
	fmt.Println("elements of my Set[T]", mySet.Elements())

}


