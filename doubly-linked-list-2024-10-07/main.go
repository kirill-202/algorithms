package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	prev  *Node
	value int
}

type DLlist struct {
	Head   *Node
	Length int
}

func InitList(val int) *DLlist {
	head := &Node{value: val}
	return &DLlist{
		Head:   head,
		Length: 1,
	}
}

func (l *DLlist) InsertNode(val int) {
    newNode := &Node{value: val, next: l.Head}
    l.Head.prev = newNode
    l.Head = newNode
    l.Length++
}

func (l *DLlist) RemoveNodeByValue(val int) (int, error) {
	n := l.Head
	for n != nil {
		if n.value == val {
			if n.prev != nil {
				n.prev.next = n.next
			} else {
				l.Head = n.next
			}
			if n.next != nil {
				n.next.prev = n.prev
			}
			l.Length--
			return n.value, nil
		}
		n = n.next
	}
	return 0, fmt.Errorf("value %d not found in the list", val)
}

func (l *DLlist) PrintList() {
	n := l.Head
	for n != nil {
		fmt.Printf("%d ", n.value)
		n = n.next
	}
	fmt.Println()
}

func main() {
	fmt.Println("Program is running...")

	firstNode := InitList(30)


	for i := 0; i < 10; i++ {
		firstNode.InsertNode(i)
	}

	firstNode.PrintList() 

	removedVal, err := firstNode.RemoveNodeByValue(5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Removed value:", removedVal)
		firstNode.PrintList() 
	}
}
