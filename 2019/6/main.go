package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	mycmd "2022-day6/cmd"

	"github.com/spf13/cobra"
)

var inputFile string
var inputText string

var rootCmd = &cobra.Command{
	Use:   "day6",
	Short: "AdventOfCode 2019 day 6",
	Long:  `AdventOfCode 2019 day 6`,
	Run: func(cmd *cobra.Command, args []string) {
		var inputs []string

		if inputFile != "" && inputText != "" {
			fmt.Println("Both 'file' and 'input' parameters were provided. Please provide only one.")
			return
		}

		if inputFile == "" && inputText == "" {
			fmt.Println("Neither 'file' nor 'input' parameter was provided. Please provide one.")
			return
		}

		if inputFile != "" {
			file, err := os.Open(inputFile)
			if err != nil {
				fmt.Println("Could not open file:", err)
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				inputs = append(inputs, scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading file:", err)
			}
		}

		if inputText != "" {
			inputs = strings.Split(inputText, "\n")
		}

		// Now you can process the inputs
		fmt.Println(mycmd.Root(inputs))
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputFile, "file", "f", "", "Input file")
	rootCmd.PersistentFlags().StringVarP(&inputText, "input", "i", "", "Input text")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
