package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

var backpacks = make([]string, 0)
var trioBackpacks = make([][]string, 0)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)

	lines := strings.Split(input, "\n")

	backpacks = append(backpacks, lines...)
	for i := 0; i < len(backpacks); i += 3 {
		trioBackpacks = append(trioBackpacks, backpacks[i:i+3])
	}

	partOne()
	partTwo()
}

func partOne() {
	total := 0

	for _, backpack := range backpacks {
		half := len(backpack) / 2
		firstHalf, secondHalf := backpack[:half], backpack[half:]
		firstHalfLetters := make(map[rune]bool)
		for _, c := range firstHalf {
			firstHalfLetters[c] = true
		}

		for _, c := range secondHalf {
			if firstHalfLetters[c] {
				total += priority(c)
				break
			}
		}
	}

	fmt.Println(total)
}

func partTwo() {
	total := 0

	for _, trio := range trioBackpacks {
		letterCounts := make(map[rune]int)
		for i, backpack := range trio {
			for _, c := range backpack {
				if letterCounts[c] == i {
					letterCounts[c]++
				}
			}
		}

		for c, count := range letterCounts {
			if count == 3 {
				total += priority(c)
				break
			}
		}
	}

	fmt.Println(total)
}

func priority(char rune) int {
	if unicode.IsUpper(char) {
		return int(char-'A') + 27
	}

	return int(char-'a') + 1
}
