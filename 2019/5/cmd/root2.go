/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var id int
var input []int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "2019-day5",
	Short: "AdventOfCode day5",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		input, _ = cmd.Flags().GetIntSlice("input")
		fmt.Printf("The input parameter is %v\n", input)
		fmt.Printf("The id parameter is %d\n", id)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute2() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("oops")
		os.Exit(1)
	}
	main()
}

func init() {
	rootCmd.Flags().IntVarP(&id, "id", "i", 0, "The digit parameter value")
	rootCmd.Flags().IntSliceP("input", "n", []int{}, "Program input")
	//rootCmd.MarkFlagRequired("id")

}

// separate each digit of a number which has four of them
func separateDigits(n int) (int, int, int, int) {
	a := n / 10000
	b := (n - a*10000) / 1000
	c := (n - a*10000 - b*1000) / 100
	d := (n - a*10000 - b*1000 - c*100)
	return a, b, c, d
}

// test the separateDigits function
func testSeparateDigits() {
	a, b, c, d := separateDigits(1002)
	fmt.Println(a, b, c, d)
}

func main() {
	testSeparateDigits()
}
