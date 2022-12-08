package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input string
var grid = make([][]int, 0)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = string(i)
	generateGrid()

	partOne, partTwo := checkTrees()
	fmt.Println(partOne)
	fmt.Println(partTwo)
}

func checkTrees() (int, int) {
	totalVisible := 0
	biggestScore := 0
	rows := len(grid)
	cols := len(grid[0])
	for rowIndex := 0; rowIndex < rows; rowIndex++ {
		for colIndex := 0; colIndex < len(grid[rowIndex]); colIndex++ {
			if rowIndex == 0 || colIndex == 0 || rowIndex == rows-1 || colIndex == cols-1 {
				totalVisible++
				continue
			}
			current := grid[rowIndex][colIndex]

			scoreTop := 1
			treeVisible := false
			for r := rowIndex - 1; 0 <= rowIndex; r-- {
				if r == 0 {
					if !treeVisible {
						treeVisible = grid[r][colIndex] < current
					}
					break
				}
				if grid[r][colIndex] >= current {
					break
				}
				scoreTop++
			}

			scoreBottom := 1
			for r := rowIndex + 1; r <= rows; r++ {
				if r == rows-1 {
					if !treeVisible {
						treeVisible = grid[r][colIndex] < current
					}
					break
				}
				if grid[r][colIndex] >= current {
					break
				}
				scoreBottom++
			}

			scoreLeft := 1
			for c := colIndex - 1; 0 <= c; c-- {
				if c == 0 {
					if !treeVisible {
						treeVisible = grid[rowIndex][c] < current
					}
					break
				}
				if grid[rowIndex][c] >= current {
					break
				}
				scoreLeft++
			}

			scoreRight := 1
			for c := colIndex + 1; c <= cols; c++ {
				if c == cols-1 {
					if !treeVisible {
						treeVisible = grid[rowIndex][c] < current
					}
					break
				}
				if grid[rowIndex][c] >= current {
					break
				}
				scoreRight++
			}

			total := scoreTop * scoreBottom * scoreLeft * scoreRight
			if total > biggestScore {
				biggestScore = total
			}

			if treeVisible {
				totalVisible++
			}

		}
	}
	return totalVisible, biggestScore
}

func generateGrid() {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row := strings.Split(line, "")
		rowConverted := make([]int, 0)
		for _, item := range row {
			itemConverted, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			rowConverted = append(rowConverted, itemConverted)
		}
		grid = append(grid, rowConverted)
	}
}
