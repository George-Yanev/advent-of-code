package main

import (
	"fmt"
	"strconv"
)

const start int = 109165
const end int = 576723

// const end int = 123789

func nextIncreasing(n int) int {
	//fmt.Println("Next number to increase: ", n)
	var closestIncreasing string
	s := strconv.Itoa(n)
	var m = map[bool]int{}
	for i, d := range s {
		if i == 0 {
			closestIncreasing += string(d)
			continue
		}

		if i, ok := m[true]; ok {
			closestIncreasing += string(s[i])
			continue
		}

		p := closestIncreasing[i-1]
		if int(d) < int(p) {
			closestIncreasing += string(p)
			m[true] = i - 1
		} else {
			closestIncreasing += string(d)
		}
	}
	c, err := strconv.Atoi(closestIncreasing)
	if err != nil {
		fmt.Printf("Not able to convert int to string: %v", closestIncreasing)
	}
	return c
}

func checkDoubles(n int) bool {
	s := strconv.Itoa(n)
	for i, d := range s {
		if i == 0 {
			continue
		}
		if int(s[i-1]) == int(d) {
			return true
		}
	}
	return false
}

func main() {
	var passwords []int
	loopBreaker := true
	currentStart := start
	for loopBreaker {
		for i := currentStart; i <= end; i++ {
			n := nextIncreasing(i)
			//fmt.Println("Current number: ", n)
			if n >= end {
				loopBreaker = false
			}

			if i != n {
				currentStart = n
				break
			}
			if checkDoubles(n) {
				passwords = append(passwords, n)
			}
		}
	}
	fmt.Printf("Number of possible passwords: %d\n", len(passwords))
	//fmt.Printf("List the passwords: %v", passwords)
}
