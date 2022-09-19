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

// contains tells whether a contains x.
func contains(a []snakePoint, x snakePoint) (int, error) {
	for i, n := range a {
		if x == n {
			return i, nil
		}
	}
	return 0, errors.New("error")
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

func calculateManhattanDistance(s1, s2 []snakePoint) (int, error) {
	intersect, err := intersection(s1, s2)
	if err != nil {
		fmt.Errorf("Error while creating the intersection - %v", err)
	}

	d, err := manhattanDistance(intersect)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return d, nil
}

func pointPathDistance(s snakePoint, s1, s2 []snakePoint) int {
	d1, err := contains(s1, s)
	if err != nil {
		fmt.Errorf("error calculating snake1 distance: %v", err)
	}

	d2, err := contains(s2, s)
	if err != nil {
		fmt.Errorf("error calculating snake2 distance: %v", err)
	}
	return d1 + d2
}

func calculateShortestPathDistance(s1, s2 []snakePoint) (int, error) {
	intersect, err := intersection(s1, s2)
	if err != nil {
		fmt.Errorf("Error while creating the intersection - %v", err)
	}

	pd := math.MaxInt
	for _, p := range intersect {
		if cd := pointPathDistance(p, s1, s2); cd < pd {
			pd = cd
		}
	}
	if pd == math.MaxInt {
		return pd, errors.New("No ShortestPathDistnace calculated")
	}
	return pd, nil
}

func main() {
	f, err := os.Open("input1")
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

	//md, err := calculateManhattanDistance(snake1, snake2)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("Manhattan Distance is: %d", md)

	pd, err := calculateShortestPathDistance(snake1, snake2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Shorted Point Distance is: %d", pd)

}
