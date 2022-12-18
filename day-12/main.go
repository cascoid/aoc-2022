package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) add(b Coordinate) Coordinate {
	return Coordinate{c.X + b.X, c.Y + b.Y}
}

var Directions = []Coordinate{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var heightMap = make(map[Coordinate]rune)

var startPoint = Coordinate{0, 0}
var endPoint = Coordinate{0, 0}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	y := 0
	for scanner.Scan() {
		for x, s := range scanner.Text() {
			heightMap[Coordinate{x, y}] = s
			if s == 'S' {
				startPoint = Coordinate{x, y}
				heightMap[Coordinate{x, y}] = 'a'
			}
			if s == 'E' {
				endPoint = Coordinate{x, y}
				heightMap[Coordinate{x, y}] = 'z'
			}
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	partOne, partTwo := BFS()
	fmt.Println(partOne)
	fmt.Println(partTwo)
}

func BFS() (int, int) {
	distance := map[Coordinate]int{endPoint: 0}
	visited := map[Coordinate]bool{endPoint: true} // block off the end point
	queue := []Coordinate{endPoint}
	var shortest *Coordinate

	for len(queue) > 0 {
		j := len(queue)
		for i := 0; i < j; i++ {
			currentCoords := queue[0]
			queue = queue[1:]

			if heightMap[currentCoords] == 'a' && shortest == nil {
				shortest = &currentCoords
			}

			for _, dir := range Directions {
				movedCoordinates := currentCoords.add(dir)

				if visited[movedCoordinates] {
					continue
				}

				_, canVisit := heightMap[movedCoordinates]

				if !canVisit {
					continue
				}

				if heightMap[currentCoords] <= heightMap[movedCoordinates]+1 {
					visited[movedCoordinates] = true
					distance[movedCoordinates] = distance[currentCoords] + 1
					queue = append(queue, movedCoordinates)
				}
			}
		}
	}

	fmt.Println(distance)

	return distance[startPoint], distance[*shortest]
}
