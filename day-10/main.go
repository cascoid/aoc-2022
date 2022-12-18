package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Operation string
	Amount    int
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
		instr := Instruction{Operation: s[0]}
		if len(s) == 2 {
			n, err := strconv.Atoi(s[1])
			if err != nil {
				panic(err)
			}
			instr.Amount = n
		}
		instructions = append(instructions, instr)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	partOne()
	partTwo()
}

func partOne() {
	res := getResult(20, 60, 100, 140, 180, 220)
	fmt.Println(res)
}

func partTwo() {
	res := getResult(240)
	fmt.Println(res)
}

func getResult(checkPoints ...int) int {
	checkPointValues := make(map[int]int)
	rows := [][]string{}
	for _, checkPoint := range checkPoints {
		checkPointValues[checkPoint] = 0
		if checkPoint/40 >= len(rows) {
			for i := len(rows); i <= checkPoint/40; i++ {
				rows = append(rows, []string{})
			}
		}
	}

	next := []int{}

	skip := false

	x := 1
	cycle := 0
	for i := 0; i < len(instructions); {
		generateLetters(&rows, cycle, x)

		cycle++
		if _, ok := checkPointValues[cycle]; ok {
			checkPointValues[cycle] = cycle * x
		}

		if !skip {
			switch instructions[i].Operation {
			case "noop":
				break
			case "addx":
				next = append(next, []int{0, instructions[i].Amount}...)
				skip = true
			}
			i++
		} else {
			skip = false
		}

		if len(next) > 0 {
			x += next[0]
			next = next[1:]
		}
	}

	for _, row := range rows {
		fmt.Println(strings.Join(row, ""))
	}

	result := 0
	for _, checkPoint := range checkPoints {
		result += checkPointValues[checkPoint]
	}

	return result
}

func generateLetters(rows *[][]string, cycle, x int) {
	col := cycle % 40
	if (cycle / 40) >= len(*rows) {
		return
	}
	row := (*rows)[cycle/40]
	if col >= x-1 && col <= x+1 {
		row = append(row, "#")
	} else {
		row = append(row, ".")
	}
	(*rows)[cycle/40] = row
}
