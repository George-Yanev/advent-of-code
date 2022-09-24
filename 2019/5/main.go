package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

const maxNoun int = 12
const maxVerb int = 2

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func ReadFile(f string) []int {
	r, err := os.Open(f)
	if err != nil {
		fmt.Println("error reading the file", err)
	}
	defer r.Close()

	reader := csv.NewReader(r)
	var input []string
	var final []int
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
		}
		input = record
	}

	for _, v := range input {
		i, _ := strconv.Atoi(v)
		final = append(final, i)
	}
	return final
}

func main() {
	for {
		input := ReadFile("input")
		fmt.Println("Input length", len(input))
		i := 0
		for i < len(input) {
			var opcode, pam1, input1, pam2, input2, output int
			pam1 = int(input[i]/100) % 10
			pam2 = int(input[i]/1000) % 10
			opcode = input[i] % 100
			fmt.Println("opcode is: ", input[i])
			fmt.Println("pam2-1 are: ", pam2, pam1)
			switch {
			case opcode == 1 || opcode == 2:
				if pam1 == 0 {
					inputTmp := input[i+1]
					input1 = input[inputTmp]
				} else {
					input1 = input[i+1]
				}
				if pam2 == 0 {
					inputTmp := input[i+2]
					input2 = input[inputTmp]
				} else {
					input2 = input[i+2]
				}
				output = input[i+3]
				if opcode == 1 {
					//fmt.Printf("For opcode 1, write to position %d = %d\n", output, input1+input2)
					input[output] = input1 + input2
				} else {
					//fmt.Printf("For opcode 2, write to position %d = %d\n", output, input1*input2)
					input[output] = input1 * input2
				}
				i += 4
			case opcode == 3:
				input1 = 1
				output = input[i+1]
				fmt.Printf("opcode 3, write to position %d = %d\n", output, input1)
				input[output] = input1
				i += 2
			case opcode == 4:
				if pam1 == 0 {
					outputTmp := input[i+1]
					output = input[outputTmp]
					fmt.Printf("Opcode 4 detected. element %d has position %d = %d\n", i, outputTmp, output)
				} else {
					output = input[i+1]
					fmt.Printf("Opcode 4 detected: elemet %d has value %d\n", i, output)
				}
				i += 2
			case opcode == 99:
				//fmt.Println("reached opcode 99...Print left-overs: ", input[i:])
				//fmt.Println("result is: ", input[0])
				return
			}
			if i > len(input) {
				break
			}
			//fmt.Println("i is: ", i)
		}
	}
}
