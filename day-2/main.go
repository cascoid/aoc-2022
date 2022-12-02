package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var values = map[string]int{
	"A": 1, // Rock
	"B": 2, // Paper
	"C": 3, // Scissors
}

var partOneResponses = map[string]string{
	"X": "A", // Rock
	"Y": "B", // Paper
	"Z": "C", // Scissors
}

var outcomes = map[string]int{
	"AA": 3,
	"AB": 6,
	"AC": 0,
	"BA": 0,
	"BB": 3,
	"BC": 6,
	"CA": 6,
	"CB": 0,
	"CC": 3,
}

var winningWeapons = map[string]string{
	"A": "B",
	"B": "C",
	"C": "A",
}

var losingWeapons = map[string]string{
	"A": "C",
	"B": "A",
	"C": "B",
}

var allRounds = make([][]string, 0)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		split := strings.Split(line, " ")
		allRounds = append(allRounds, split)
	}

	partOne()

	partTwo()
}

func partOne() {
	myTotal := 0
	for _, round := range allRounds {
		myWeapon := partOneResponses[round[1]]
		myTotal += outcomes[round[0]+myWeapon] + values[myWeapon]
	}
	fmt.Println(myTotal)
}

func partTwo() {
	myTotal := 0
	for _, round := range allRounds {
		myWeapon := "A"
		switch round[1] {
		case "X":
			myWeapon = losingWeapons[round[0]]
		case "Y":
			myWeapon = round[0]
		case "Z":
			myWeapon = winningWeapons[round[0]]
		}
		myTotal += outcomes[round[0]+myWeapon] + values[myWeapon]
	}
	fmt.Println(myTotal)
}
