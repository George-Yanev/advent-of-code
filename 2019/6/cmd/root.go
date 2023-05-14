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

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) AddChild(name string) {
	child := &Node{Name: name}
	n.Children = append(n.Children, child)
}

func (n *Node) printDFS(indent int) {
	fmt.Println(strings.Repeat(" ", indent), n.Name)

	for _, child := range n.Children {
		child.printDFS(indent)
	}
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
	fmt.Printf("The in parameter is:\n%v\n", in)
	// loop over the input
	// the input is a two strings separated by a ) character
	// the first string is the body and the second string is the orbiter
	tree := make(map[string]*Node)
	for _, line := range strings.Split(in, "\n") {

		// split the line into the body and the orbiter
		l := strings.Split(line, ")")
		body, orbiter := l[0], l[1]

		_, ok := tree[body]
		if !ok {
			tree[body] = &Node{Name: body}
		}

		tree[body].AddChild(orbiter)

	}

	fmt.Println("start printing")
	tree["COM"].printDFS(2)
	fmt.Printf("tree is %v\n", tree["B"])

}

func main() {
	execute(input)
}
