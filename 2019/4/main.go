package main

import (
	"fmt"
	"strconv"
)

const start int = 109165
// const end int = 576723
const end int = 111222

func nextIncreasing(n int) int {
	var closestIncreasing string
	s := strconv.Itoa(n)
	for i, d := range s {
		if i == 0 {
			closestIncreasing += string(d)
			continue
		}

		p := int(closestIncreasing[i-1])
		if int(d) < p {
			closestIncreasing += string(p)
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

func main() {
	var passwords []int
	loopBreaker := true
	currentStart := start
	for loopBreaker {
		for i := currentStart; i <= end; i++ {
			n := nextIncreasing(i)
			f n >= end {
				loopBreaker = false
			}

			if i != n {
				currentStart = n
				break
			}
			passwords = append(passwords, n)
		}
	}
	fmt.Printf("Number of possible passwords: %d\n", len(passwords))
	fmt.Printf("List the passwords: %v", passwords)
}
