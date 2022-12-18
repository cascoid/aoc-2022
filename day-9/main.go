package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	Y int
	X int
}

func (c *Coordinate) add(b Coordinate) {
	c.X += b.X
	c.Y += b.Y
}

type Instruction struct {
	Direction Coordinate
	Amount    int
}

var Directions = map[string]Coordinate{
	"L": {0, -1},
	"R": {0, 1},
	"U": {1, 0},
	"D": {-1, 0},
	"":  {0, 0},
}

var instructions []Instruction

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		n, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, Instruction{Directions[s[0]], n})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	partOne()
	partTwo()
}

func partOne() {
	head := Coordinate{0, 0}
	tail := Coordinate{0, 0}

	var visited = map[Coordinate]bool{}

	for _, instruction := range instructions {
		moveRope([]*Coordinate{&head, &tail}, instruction, visited)
	}

	fmt.Println(len(visited))
}

func partTwo() {
	segments := []*Coordinate{}
	var visited = map[Coordinate]bool{}
	for i := 0; i < 10; i++ {
		segments = append(segments, &Coordinate{0, 0})
	}

	for _, instruction := range instructions {
		moveRope(segments, instruction, visited)
	}

	fmt.Println(len(visited))
}

func moveRope(segments []*Coordinate, instruction Instruction, visited map[Coordinate]bool) {
	head := segments[0]
	for i := 0; i < instruction.Amount; i++ {
		head.X += instruction.Direction.X
		head.Y += instruction.Direction.Y

		for i := 1; i < len(segments); i++ {
			first, second := moveTail(segments[i-1], segments[i])
			segments[i].add(Directions[first])
			segments[i].add(Directions[second])

			if i == len(segments)-1 {
				visited[*segments[i]] = true
			}
		}

	}

}

func moveTail(head, tail *Coordinate) (string, string) {
	if head.X == tail.X+2 && head.Y == tail.Y {
		return "R", ""
	}
	if head.X == tail.X-2 && head.Y == tail.Y {
		return "L", ""
	}
	if head.Y == tail.Y+2 && head.X == tail.X {
		return "U", ""
	}
	if head.Y == tail.Y-2 && head.X == tail.X {
		return "D", ""
	}
	if head.X == tail.X+2 && head.Y == tail.Y+1 || head.X == tail.X+1 && head.Y == tail.Y+2 ||
		head.X == tail.X+2 && head.Y == tail.Y+2 {
		return "R", "U"
	}
	if head.X == tail.X+2 && head.Y == tail.Y-1 || head.X == tail.X+1 && head.Y == tail.Y-2 ||
		head.X == tail.X+2 && head.Y == tail.Y-2 {
		return "R", "D"
	}
	if head.X == tail.X-2 && head.Y == tail.Y+1 || head.X == tail.X-1 && head.Y == tail.Y+2 ||
		head.X == tail.X-2 && head.Y == tail.Y+2 {
		return "L", "U"
	}
	if head.X == tail.X-2 && head.Y == tail.Y-1 || head.X == tail.X-1 && head.Y == tail.Y-2 ||
		head.X == tail.X-2 && head.Y == tail.Y-2 {
		return "L", "D"
	}

	return "", ""
}
