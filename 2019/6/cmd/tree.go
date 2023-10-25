package cmd

import (
	"fmt"
	"strings"
)

// type Nodes []*Node

type Node struct {
	Name   string
	Parent string
	Left   string
	Right  string
}

type HashMapTree struct {
	hash map[string]*Node
}

type Queue struct {
	nodes []*Node
}

func (q *Queue) Enqueue(n *Node) {
	q.nodes = append(q.nodes, n)
}

func (q *Queue) Dequeue() *Node {
	if len(q.nodes) == 0 {
		return nil
	}

	item := q.nodes[0]
	q.nodes = q.nodes[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}

// String function to print all Node properties
func (n *Node) String() string {
	// return fmt.Sprintf("Node %v has %v children and parent %v", n.Name, len(n.Children), n.Parent.Name)
	return fmt.Sprintf(n.Name)
}

func ConvertInputToMap(m []string) map[string][]string {
	myMap := make(map[string][]string)
	for _, value := range m {
		t := strings.Split(value, ")")
		lcom := t[0]
		orbiter := t[1]
		if _, ok := myMap[lcom]; !ok {
			myMap[lcom] = []string{orbiter}
		} else {
			myMap[lcom] = append(myMap[lcom], orbiter)
		}

	}
	return myMap
}

func (b *HashMapTree) Insert(parentName string, nodeName string, childrensNames []string, initialMap map[string][]string) {
	if len(childrensNames) == 0 {
		return
	}

	left := childrensNames[0]
	node := &Node{
		Name:   nodeName,
		Parent: parentName,
		Left:   left,
		Right:  "",
	}
	b.Insert(nodeName, left, initialMap[left], initialMap)

	if len(childrensNames) > 1 {
		right := childrensNames[1]
		node.Right = right
		b.Insert(nodeName, right, initialMap[right], initialMap)
	}

	b.hash[nodeName] = node

}

func (b *HashMapTree) Print(nodeName string, level int) {
	node, ok := b.hash[nodeName]
	if !ok {
		return
	}

	for i := 0; i < level; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("Node: %s, Parent: %s, Left child: %s, Right child: %s\n", node.Name, node.Parent, node.Left, node.Right)
	b.Print(node.Left, level+1)
	b.Print(node.Right, level+1)
}
