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
var in []int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "2019-day5",
	Short: "AdventOfCode day5",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		in, _ = cmd.Flags().GetIntSlice("input")
		fmt.Printf("The in parameter is %v\n", in)
		fmt.Printf("The id parameter is %d\n", id)
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
	a, b, c, d := separateDigits(1003)
	fmt.Println(a, b, c, d)
}

// based on testSeparateDigits function first parameter is the opcode
// second parameter is the mode of the first parameter
// third parameter is the mode of the second parameter
// fourth parameter is the mode of the third parameter
func getOpcode(n int) (int, int, int, int) {
	a, b, c, d := separateDigits(n)
	return d, c, b, a
}

// based on opcode it will be decided what to do
// 1 - add
// 2 - multiply
// 3 - input
// 4 - output
// 5 - jump-if-true
// 6 - jump-if-false
// 7 - less than
// 8 - equals
// 99 - halt
func getOperation(n int) int {
	a, _, _, _ := getOpcode(n)
	return a
}

// use the opcode to to get all necessary parameters
// for the opcode operation
func getParameterModes(n int) (int, int, int) {
	_, c, b, _ := getOpcode(n)
	return c, b, 0
}

// based on the mode it will be decided what to do
// 0 - position mode
// 1 - immediate mode
// If the parameter is in position mode, its value is the value stored at the address given by the parameter.
// If the parameter is in immediate mode, its value is simply the value of the parameter.
func getParameter(n int, mode int) int {
	if mode == 0 {
		return in[n]
	}
	return n
}

// opcode 1 - add
func add(a, b int) int {
	return a + b
}

// opcode 2 - multiply
func multiply(a, b int) int {
	return a * b
}

// opcode 3 - input
func input() int {
	var a int
	fmt.Println("Please enter a number: ")
	fmt.Scan(&a)
	return a
}

// opcode 4 - output
func output(a int) {
	fmt.Println("Output is: ", a)
}

// opcode 5 - jump-if-true
func jumpIfTrue(a, b int) int {
	if a != 0 {
		return b
	}
	return 0
}

// opcode 6 - jump-if-false
func jumpIfFalse(a, b int) int {
	if a == 0 {
		return b
	}
	return 0
}

// opcode 7 - less than
func lessThan(a, b int) int {
	if a < b {
		return 1
	}
	return 0
}

// opcode 8 - equals
func equals(a, b int) int {
	if a == b {
		return 1
	}
	return 0
}

// loop through the input and execute the operations
func execute(in []int, id int) {
	for i := 0; i < len(in); {
		opcode := getOperation(in[i])
		fmt.Println("opcode is: ", opcode)
		switch {
		case opcode == 1:
			// add
			// get parameter modes
			mode1, mode2, _ := getParameterModes(in[i])
			// get the values
			parameter1 := getParameter(in[i+1], mode1)
			parameter2 := getParameter(in[i+2], mode2)
			// calculate the result
			result := add(parameter1, parameter2)
			// store the result
			in[in[i+3]] = result
			// increment the index
			i += 4
		case opcode == 2:
			// multiply
			// get parameters
			mode1, mode2, _ := getParameterModes(in[i])
			// get the values
			parameter1 := getParameter(in[i+1], mode1)
			parameter2 := getParameter(in[i+2], mode2)
			// calculate the result
			result := multiply(parameter1, parameter2)
			// store the result
			in[in[i+3]] = result
			// increment the index
			i += 4
		case opcode == 3:
			// input
			// get the value
			parameter1 := id
			// store the result
			in[in[i+1]] = parameter1
			// increment the index
			i += 2
		case opcode == 4:
			// output
			// get the value
			parameter1 := in[in[i+1]]
			// output the result
			output(parameter1)
			// increment the index
			i += 2
		case opcode == 5:
			mode1, mode2, _ := getParameterModes(in[i])
			parameter1 := getParameter(in[i+1], mode1)
			parameter2 := getParameter(in[i+2], mode2)
			result := jumpIfTrue(parameter1, parameter2)
			if result == parameter2 {
				i = result
			} else {
				i += 3
			}
		case opcode == 6:
			mode1, mode2, _ := getParameterModes(in[i])
			parameter1 := getParameter(in[i+1], mode1)
			parameter2 := getParameter(in[i+2], mode2)
			result := jumpIfFalse(parameter1, parameter2)
			if result == parameter2 {
				i = result
			} else {

				i += 3
			}
		case opcode == 7:
			mode1, mode2, _ := getParameterModes(in[i])
			parameter1 := getParameter(in[i+1], mode1)
			parameter2 := getParameter(in[i+2], mode2)
			if parameter1 < parameter2 {
				in[in[i+3]] = 1
			} else {
				in[in[i+3]] = 0
			}
			i += 4
		case opcode == 8:
			mode1, mode2, _ := getParameterModes(in[i])
			parameter1 := getParameter(in[i+1], mode1)
			parameter2 := getParameter(in[i+2], mode2)
			if parameter1 == parameter2 {
				in[in[i+3]] = 1

			} else {
				in[in[i+3]] = 0
			}
			i += 4
		case opcode == 99:
			// halt
			fmt.Println("Halt. Program finished. Result is - ", in)
			return
		default:
			fmt.Println("Unknown opcode")
			return
		}
	}
}

func main() {
	execute(in, id)
}
