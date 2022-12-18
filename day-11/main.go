package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	Items            []int
	TotalInspections int
	Operation        Operation
	DivisibleTest    int
	TrueMonkey       int
	FalseMonkey      int
}

type Operation struct {
	Operator     string
	Number       int
	NumberIsSelf bool
}

var monkeys = []*Monkey{{}}
var partTwoMonkeys = []*Monkey{{}}

var divisor int = 1

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	monkeyInput := [][]string{{}}

	currentMonkey := 0
	for _, line := range lines {
		if line == "" {
			monkeys = append(monkeys, &Monkey{})
			partTwoMonkeys = append(partTwoMonkeys, &Monkey{})
			monkeyInput = append(monkeyInput, []string{})
			currentMonkey++
			continue
		}
		monkeyInput[currentMonkey] = append(monkeyInput[currentMonkey], line)
	}

	for i, monkey := range monkeys {
		monkey.parseMonkey(monkeyInput[i])
	}

	for i, monkey := range partTwoMonkeys {
		monkey.parseMonkey(monkeyInput[i])
		divisor *= monkey.DivisibleTest
	}

	partOne(monkeys)
	partTwo(partTwoMonkeys)
}

func partOne(monkeys []*Monkey) {
	for i := 0; i < 20; i++ {
		completeRound(monkeys, true)
	}

	fmt.Println(getHighestItems(monkeys))
}

func partTwo(monkeys []*Monkey) {
	for i := 0; i < 10000; i++ {
		completeRound(monkeys, false)
	}

	fmt.Println(getHighestItems(monkeys))
}

func getHighestItems(monkeys []*Monkey) int {
	highestItems := []int{0, 0}
	for _, monkey := range monkeys {
		if monkey.TotalInspections > highestItems[0] {
			highestItems[1] = highestItems[0]
			highestItems[0] = monkey.TotalInspections
		} else if monkey.TotalInspections > highestItems[1] {
			highestItems[1] = monkey.TotalInspections
		}
	}
	return highestItems[0] * highestItems[1]
}

func (monkey *Monkey) parseMonkey(input []string) {
	for i, line := range input {
		switch i {
		case 1:
			monkey.Items = splitItems(line)
		case 2:
			monkey.Operation = getOperation(line)
		case 3:
			monkey.DivisibleTest = getDivisibleTest(line)
		case 4:
			monkey.TrueMonkey = getMonkeyThrow(line)
		case 5:
			monkey.FalseMonkey = getMonkeyThrow(line)
		}
	}
}

func completeRound(monkeys []*Monkey, divideByThree bool) {
	for i, monkey := range monkeys {
		for _, item := range monkey.Items {

			newWorry := item
			secondNumber := monkey.Operation.Number
			if monkey.Operation.NumberIsSelf {
				secondNumber = item
			}
			if monkey.Operation.Operator == "*" {
				newWorry = item * secondNumber
			} else {
				newWorry = item + secondNumber
			}

			if divideByThree {
				newWorry = newWorry / 3
			} else {
				newWorry = newWorry % divisor
			}
			if newWorry%monkey.DivisibleTest == 0 {
				monkeys[monkey.TrueMonkey].Items = append(monkeys[monkey.TrueMonkey].Items, newWorry)
			} else {
				monkeys[monkey.FalseMonkey].Items = append(monkeys[monkey.FalseMonkey].Items, newWorry)
			}
			monkey.TotalInspections++
		}
		monkeys[i].Items = []int{}
	}
}

func splitItems(input string) []int {
	input = strings.Trim(input, " ")
	input = strings.TrimPrefix(input, "Starting items: ")
	input = strings.ReplaceAll(input, " ", "")
	s := strings.Split(input, ",")
	items := make([]int, len(s))
	for i, item := range s {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		items[i] = num
	}
	return items
}

func getOperation(input string) Operation {
	input = strings.Trim(input, " ")
	input = strings.TrimPrefix(input, "Operation: new = ")
	split := strings.Split(input, " ")
	if len(split) == 3 {
		op := Operation{Operator: split[1]}
		if split[0] == split[2] {
			op.NumberIsSelf = true
			op.Number = 0
			return op
		}

		num, err := strconv.Atoi(split[2])
		if err != nil {
			panic(err)
		}
		op.Number = num
		return op
	} else {
		panic("Invalid operation")
	}
}

func getDivisibleTest(input string) int {
	input = strings.Trim(input, " ")
	input = strings.TrimPrefix(input, "Test: divisible by ")
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return num
}

func getMonkeyThrow(input string) int {
	input = strings.Trim(input, " ")
	input = strings.TrimPrefix(input, "If true: throw to monkey ")
	input = strings.TrimPrefix(input, "If false: throw to monkey ")
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return num
}
