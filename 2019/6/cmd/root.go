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

var in []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "2019-day5",
	Short: "AdventOfCode day5",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		in, _ = cmd.Flags().GetStringSlice("input")
		fmt.Printf("The in parameter is %v\n", in)
	},
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
	rootCmd.Flags().IntSliceP("input", "n", []int{}, "Program input")
	rootCmd.MarkFlagRequired("input")

}

func execute(in []string) {
	fmt.Printf("The in parameter is %v\n", in)

	// initiate a map of string to int
	structure := make(map[string]int)

	// loop over the input
	// the input is a two strings separated by a ) character
	// the first string is the body and the second string is the orbiter
	for i := 1; i <len(in); i++ {
		// split the line into the body and the orbiter
		l := strings.Split(line, ")")
		body := l[0]
		orbiter := l[1]
		// add the orbiter to the body's list of orbiters
		structure[orbiter] = 
	}
}

func main() {
	execute(in)
}
