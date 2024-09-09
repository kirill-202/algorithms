package main

import (
	"fmt"
	"os"
	"slices"

)

/*

Description:

Implement a queue data structure using two stacks. The queue should have the following operations:

Enqueue(x int): Add an element to the end of the queue.
Dequeue() int: Remove the element from the front of the queue.
Peek() int: Get the element at the front of the queue without removing it.
IsEmpty() bool: Check if the queue is empty.
Requirements:

Do not use any built-in queue or list methods that directly provide the result.
You must only use stacks (LIFO data structure) for implementing the queue functionality.
Design your solution with efficient time complexity.

*/



type Stack struct {
	elements []int
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
	fmt.Printf("%d has been added to %v\n", value, s.elements)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("the slice is empty, cannot pop")
	}

	pop_value := s.elements[len(s.elements)-1] 
	s.elements = s.elements[:len(s.elements)-1]

	fmt.Printf("%d has been removed from elements %v, current length is %v\n", pop_value, s.elements, len(s.elements))
	return pop_value, nil
} 

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0 
}


type Queue struct{
	frontStack Stack
	backStack Stack //will be reversed frontStack
}

func (q *Queue) Enqueue(value int) {
	q.backStack.Push(value)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("the queue is empty, cannot deque")
	}
	if q.frontStack.IsEmpty() {
		fmt.Println("Transferring elements from backstack...")
		if !q.backStack.IsEmpty() {
			var tmp []int
			for {
				value, err := q.backStack.Pop()
				if err != nil {
					break
				}
				tmp = append(tmp, value)
			}

			slices.Reverse(tmp)
			for _, value := range tmp {
				q.frontStack.Push(value)
			}
		}
		

		return q.frontStack.Pop()
	}
	return q.frontStack.Pop()

}


func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		fmt.Println("The queue is empty, cannot peek")
		os.Exit(1)
	}
	return q.frontStack.elements[0], nil
}

func (q *Queue) IsEmpty() bool {
	if len(q.frontStack.elements) == 0 && len(q.backStack.elements) == 0{
		return true
	}
	return false
}


func main() {
	elems := []int{1,2,3,4,5}
	stackfront := Stack{
		elements: elems,
	}

	var backfront Stack



	my_queue := Queue{
		frontStack: stackfront,
		backStack: backfront,
	}

	my_queue.Enqueue(6)
	my_queue.Enqueue(7)
	my_queue.Enqueue(8)
	fmt.Println(my_queue)
	for i:= 0; i <10; i++ {
		my_queue.Dequeue()
	}
	
}



