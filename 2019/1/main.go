package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func moduleFuelRequirement(m int) int {
	return int(m/3) - 2

}

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Errorf("err read file: %v", err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var m []string
	for fileScanner.Scan() {
		m = append(m, fileScanner.Text())
	}
	var total int
	for _, m := range m {
		i, err := strconv.Atoi(m)
		if err != nil {
			fmt.Errorf("Cannot convert this mass to int: %v", m)
		}
		f := moduleFuelRequirement(i)
		total += f
		fmt.Println("Initial fuel requirement: ", f)
		for f >= 9 {
			f = moduleFuelRequirement(f)
			fmt.Println("Each following iteration: ", f)
			total += f
		}
	}
	fmt.Println("Total mass is ", total)
}
