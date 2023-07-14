package cmd

import (
	"fmt"
	"strings"
)

type StringSlice []string

type Node struct {
	Name     string
	Children []*Node
	Parent   *Node
}

// String function to print all Node properties
func (n *Node) String() string {
	return fmt.Sprintf("Node %v has %v children and parent %v", n.Name, len(n.Children), n.Parent.Name)
}

func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

func (n *Node) printDFS(indent int) {
	fmt.Println(strings.Repeat(" ", indent), n.Name)

	for _, child := range n.Children {
		child.printDFS(indent + 1)
	}
}

var p StringSlice

// print recursively the parent Name of the Node
func (n *Node) printParents() *StringSlice {
	if n.Parent != nil {
		// fmt.Printf("%v,", n.Parent.Name)
		p = append(p, n.Parent.Name)
		n.Parent.printParents()
	}
	return &p
}

// reverse the elements of StringSlice
func (s *StringSlice) Reverse() StringSlice {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		// fmt.Printf("i=%v, j=%v\n", i, j)
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
	return *s
}

func (s StringSlice) Difference(other StringSlice) StringSlice {
	otherSet := make(map[string]bool)
	for _, str := range other {
		otherSet[str] = true
	}

	var diff StringSlice
	for _, str := range s {
		if !otherSet[str] {
			diff = append(diff, str)
		}
	}
	return diff
}

// Traverse the Tree using the Breadth-First Traversal algorithm
func (n *Node) printBFT() {
	// Create a queue (FIFO)
	queue := []*Node{n}

	for len(queue) > 0 {
		// Pop the first element from the queue
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.Name)

		// Add all children to the queue
		for _, child := range node.Children {
			queue = append(queue, child)
		}
	}
}

// create a Queue to append and remove Nodes from it
type Queue []*Node

func (q *Queue) Push(node *Node) {
	*q = append(*q, node)
}

func (q *Queue) Pop() *Node {
	if len(*q) == 0 {
		return nil
	}

	node := (*q)[0]
	*q = (*q)[1:]
	return node
}
