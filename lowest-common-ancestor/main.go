package main

import (
	"fmt"
)

/*
Given the root of a binary tree, return the lowest common ancestor (LCA) of its deepest leaves.

The lowest common ancestor is defined as the lowest node in the tree that 
has both deepest leaves as descendants (where a node can be a descendant of itself).

*/

type Node struct {
	Right *Node
	Left *Node
	Value int
}

func (n *Node) MaxDepth() int {
	if n == nil {
		return 0
	}
	leftDepth := n.Left.MaxDepth()
	rightDepth := n.Right.MaxDepth()

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func (n *Node) Insert(value int) {
	if n == nil {
		return
	} else if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func findLCADeepestLeaves(n *Node, currentDepth int, maxDepth int) (*Node, int) {
	if n == nil {
		return nil, currentDepth
	}

	if n.Left == nil && n.Right == nil {
		if currentDepth == maxDepth {
			return n, maxDepth
		}
		return nil, currentDepth
	}


	leftLCA, leftDepth := findLCADeepestLeaves(n.Left, currentDepth+1, maxDepth)
	rightLCA, rightDepth := findLCADeepestLeaves(n.Right, currentDepth+1, maxDepth)

	if leftLCA != nil && rightLCA != nil {
		return n, maxDepth
	}

	if leftLCA != nil {
		return leftLCA, leftDepth
	}
	return rightLCA, rightDepth
}

func (n *Node) FindLCAOfDeepestLeaves() *Node {
	maxDepth := n.MaxDepth()
	lca, _ := findLCADeepestLeaves(n, 1, maxDepth)
	return lca
}

func main() {
	root := &Node{Value: 10}

	values := []int{5, 15, 3, 7, 12, 18, 20, 130, 500, 4, 3, 100, 7, 9, 10, 22, 21}
	for _, v := range values {
		root.Insert(v)
	}
	fmt.Printf("The maximum depth of the tree is: %d\n", root.MaxDepth())

	lca := root.FindLCAOfDeepestLeaves()
	if lca != nil {
		fmt.Printf("The Lowest Common Ancestor of the deepest leaves is: %d\n", lca.Value)
	} else {
		fmt.Println("The tree does not have enough nodes to find an LCA.")
	}

}