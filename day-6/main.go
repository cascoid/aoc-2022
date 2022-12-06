package main

import (
	"fmt"
	"io/ioutil"
)

var input []byte

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input = i

	partOne()
	partTwo()
}

func partOne() {
	r := findUniqueCharSequence(4)
	fmt.Println(r)
}

func partTwo() {
	r := findUniqueCharSequence(14)
	fmt.Println(r)
}

func findUniqueCharSequence(markerLength int) int {
	prevChars := []byte{}
	for i := 0; i < len(input); i++ {
		r := input[i]
		index := indexIfContains(prevChars, r)
		if index != -1 {
			prevChars = prevChars[index+1:]
		}
		prevChars = append(prevChars, r)
		if len(prevChars) == markerLength {
			return i + 1
		}
	}

	return -1
}

func indexIfContains(arr []byte, elem byte) int {
	for i, a := range arr {
		if a == elem {
			return i
		}
	}
	return -1
}
