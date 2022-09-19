package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 74, 0
// 74, -29
// 157, -29
// 157, 54
type snakePoint struct {
	x, y int
}

func (s snakePoint) String() string {
	return fmt.Sprintf("snakePoint{x:%v, y:%v}", s.x, s.y)
}

func newSnakePath(snakeString string) []snakePoint {
	snakeSlice := strings.Split(snakeString, ",")
	snake := []snakePoint{{0, 0}}

	for i := 0; i < len(snakeSlice); i++ {
		steps, err := strconv.Atoi(snakeSlice[i][1:])
		if err != nil {
			fmt.Errorf("Error in the snake input: %v", steps)
		}

		p := snake[len(snake)-1]
		switch d := snakeSlice[i][0]; string(d) {
		case "L":
			for i := 1; i <= steps; i++ {
				snake = append(snake, snakePoint{p.x - i, p.y})
			}
		case "R":
			for i := 1; i <= steps; i++ {
				snake = append(snake, snakePoint{p.x + i, p.y})
			}
		case "D":
			for i := 1; i <= steps; i++ {
				snake = append(snake, snakePoint{p.x, p.y - i})
			}
		case "U":
			for i := 1; i <= steps; i++ {
				snake = append(snake, snakePoint{p.x, p.y + i})
			}
		}
	}
	return snake
}

func intersection(a, b []snakePoint) ([]snakePoint, error) {

	// uses empty struct (0 bytes) for map values.
	m := make(map[snakePoint]struct{}, len(b))

	// cached
	for _, v := range b {
		m[v] = struct{}{}
	}

	var s []snakePoint
	for _, v := range a {
		if _, ok := m[v]; ok {
			s = append(s, v)
		}
	}

	return s[1:], nil
}

func manhattanDistance(s []snakePoint) (int, error) {
	md := math.MaxInt
	for _, p := range s {
		if cd := int(math.Abs(float64(p.x)) + math.Abs(float64(p.y))); cd < md {
			md = cd
		}
	}
	if md == math.MaxInt {
		return md, errors.New("No Manhattan Distance found")
	}
	return md, nil

}

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var snakes []string
	for fileScanner.Scan() {
		snakes = append(snakes, fileScanner.Text())
	}

	snake1 := newSnakePath(snakes[0])
	snake2 := newSnakePath(snakes[1])
	intersect, err := intersection(snake1, snake2)
	if err != nil {
		fmt.Errorf("Error while creating the intersection - %v", err)
	}

	d, err := manhattanDistance(intersect)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Manhattan Distance: %v\n", d)
}
