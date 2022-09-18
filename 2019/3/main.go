package main

import (
	"fmt"
	"strconv"
	"strings"
)

const snake1Input string = "R75,D30,R83,U83,L12,D49,R71,U7,L72"
const snake2Input string = "U62,R66,U55,R34,D71,R55,D58,R83"

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

var snake []*snakePoint

func newSnakePath(snakeString string) []*snakePoint {
	snakeSlice := strings.Split(snakeString, ",")
	snake = append(snake, &snakePoint{0, 0})

	for i := 0; i < len(snakeSlice); i++ {
		steps, err := strconv.Atoi(snakeSlice[i][1:])
		if err != nil {
			fmt.Errorf("Error in the snake input: %v", steps)
		}

		p := snake[len(snake)-1]
		switch d := snakeSlice[i][0]; string(d) {
		case "L":
			snake = append(snake, &snakePoint{p.x - steps, p.y})
		case "R":
			snake = append(snake, &snakePoint{p.x + steps, p.y})
		case "D":
			snake = append(snake, &snakePoint{p.x, p.y - steps})
		case "U":
			snake = append(snake, &snakePoint{p.x, p.y + steps})
		}
	}
	return snake
}

func intersection(a, b []*snakePoint) ([]*snakePoint, error) {

	// uses empty struct (0 bytes) for map values.
	m := make(map[*snakePoint]struct{}, len(b))

	// cached
	for _, v := range b {
		m[v] = struct{}{}
	}

	var s []*snakePoint
	for _, v := range a {
		if _, ok := m[v]; ok {
			s = append(s, v)
		}
	}

	return s, nil
}

func main() {
	snake1 := newSnakePath(snake1Input)
	snake2 := newSnakePath(snake2Input)
	fmt.Printf("Snake1 is: %v\n", snake1)
	fmt.Printf("Snake2 is: %v\n", snake2)
}
