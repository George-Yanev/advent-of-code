/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var total int
var weight int

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

var input string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "2019-day6",
	Short: "AdventOfCode day6",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("oops")
		os.Exit(1)
	}
	main()
}

func init() {
	rootCmd.Flags().StringVarP(&input, "input", "n", "", "Program input")
	rootCmd.MarkFlagRequired("input")

}

func execute(in string) {
	// fmt.Printf("The in parameter is:\n%v\n", in)
	// loop over the input
	// the input is a two strings separated by a ) character
	// the first string is the body and the second string is the orbiter
	tree := make(map[string]*Node)
	for _, line := range strings.Split(in, "\n") {

		// split the line into the body and the orbiter
		l := strings.Split(line, ")")
		body, orbiter := l[0], l[1]

		_, ok := tree[body] // check if the body is already in the tree
		if !ok {
			tree[body] = &Node{Name: body}
		}

		_, ok = tree[orbiter] // check if the orbiter is already in the tree
		if !ok {
			tree[orbiter] = &Node{Name: orbiter, Parent: tree[body]}
		} else {
			ok := tree[orbiter].Parent // check if the orbiter has a parent
			if ok != nil {
				fmt.Printf("orbiter %v already has a parent %v\n", orbiter, tree[orbiter].Parent.Name)
			} else {
				tree[orbiter].Parent = tree[body]
			}
		}

		tree[body].AddChild(tree[orbiter])

	}

}

func main() {
	execute(input)
}
