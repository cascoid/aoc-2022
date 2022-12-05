package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var instructions []string

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(b), "\n\n")

	instructions = strings.Split(input[1], "\n")

	input[0] = strings.ReplaceAll(input[0], "    ", " [-]")
	input[0] = strings.ReplaceAll(input[0], "] [", "")
	input[0] = strings.ReplaceAll(input[0], "[", "")
	input[0] = strings.ReplaceAll(input[0], "]", "")

	splitInputs := strings.Split(input[0], "\n")
	splitInputs = splitInputs[:len(splitInputs)-1]
	stacks := make([][]string, len(splitInputs[0]))
	for _, splitInput := range splitInputs {
		crates := strings.Split(splitInput, "")
		for i, crate := range crates {
			if crate != "-" {
				stacks[i] = append(stacks[i], crate)
			}
		}
	}

	p2Stacks := make([][]string, len(stacks))
	copy(p2Stacks, stacks)

	partOne(stacks)
	partTwo(p2Stacks)
}

func partOne(stacks [][]string) {
	executeInstructions(stacks, false)
}

func partTwo(stacks [][]string) {
	executeInstructions(stacks, true)
}

func executeInstructions(stacks [][]string, batchMove bool) {
	for _, instruction := range instructions {
		expr := regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
		matches := expr.FindStringSubmatch(instruction)
		amount, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			panic(err)
		}
		from, err := strconv.ParseInt(matches[2], 10, 64)
		if err != nil {
			panic(err)
		}
		to, err := strconv.ParseInt(matches[3], 10, 64)
		if err != nil {
			panic(err)
		}
		stacks = moveCrates(stacks, amount, from, to, batchMove)
	}
	printResult(stacks)
}

func moveCrates(stacks [][]string, amount, from, to int64, batchMove bool) [][]string {
	cratesToMove := stacks[from-1][:amount]
	stacks[from-1] = stacks[from-1][amount:]
	if batchMove {
		// Reverse the order
		for a, b := 0, len(cratesToMove)-1; a < b; a, b = a+1, b-1 {
			cratesToMove[a], cratesToMove[b] = cratesToMove[b], cratesToMove[a]
		}
	}
	stacks[to-1] = prependInOrder(stacks[to-1], cratesToMove)
	return stacks
}

func prependInOrder(arr []string, items []string) []string {
	for _, item := range items {
		arr = append([]string{item}, arr...)
	}
	return arr
}

func printResult(stacks [][]string) {
	result := make([]string, len(stacks))
	for i, stack := range stacks {
		result[i] = stack[0]
	}
	fmt.Println(strings.Join(result, ""))
}
