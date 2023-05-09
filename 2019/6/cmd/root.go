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
	Data     interface{}
	Children []*Node
}

func (n *Node) AddChild(data interface{}) {
	child := &Node{Data: data}
	n.Children = append(n.Children, child)
}

func (n *Node) printDFS(indent int) {
	fmt.Println(strings.Repeat(" ", indent), n.Data)

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
	fmt.Printf("The in parameter is %v\n", in)
	orbits := map[string]string{}
	// loop over the input
	// the input is a two strings separated by a ) character
	// the first string is the body and the second string is the orbiter
	for _, line := range strings.Split(in, "\n") {
		// split the line into the body and the orbiter
		l := strings.Split(line, ")")
		body := l[0]
		orbiter := l[1]
		// add the orbiter to the body's list of orbiters
		orbits[body] = orbiter
	}
	fmt.Println("map structure is", orbits)

	// create the tree and the root node
	tree := &Node{Data: "COM"}
	for i := 0; i < len(orbits); i++ {

	}
	// create the tree
	for k, v := range orbits {
		fmt.Printf("key is %s, value is %s\n", k, v)

		//

	}
}

func main() {
	execute(input)
}
