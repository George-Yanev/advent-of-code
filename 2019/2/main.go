package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

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

func ReadFile(r io.Reader) []int {
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
	f, err := os.Open("input")
	if err != nil {
		fmt.Println("error reading the file", err)
	}
	defer f.Close()
	input := ReadFile(f)
	fmt.Println(len(input))

	i := 0
	for i < len(input) {
		var opcode, input1, input2, output int
		if len(input) > i+3 {
			opcode = input[i]
			input1 = input[i+1]
			input2 = input[i+2]
			output = input[i+3]
			i += 4
		} else {
			fmt.Println("end reached")
			fmt.Println(input[i:])
			opcode = input[i]
			i = len(input)
		}
		switch {
		case opcode == 1:
			input[output] = input[input1] + input[input2]
		case opcode == 2:
			input[output] = input[input1] * input[input2]
		case opcode == 99:
			fmt.Println("reached opcode 99...exit")
			fmt.Println(input)
			return
		}
	}
}
