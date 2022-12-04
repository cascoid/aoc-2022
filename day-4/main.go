package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var pairs = make([][][]int64, 0)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		stringPairs := strings.Split(line, ",")
		minFirst, maxFirst := minMaxPair(stringPairs[0])
		minSecond, maxSecond := minMaxPair(stringPairs[1])
		pairs = append(pairs, [][]int64{{minFirst, maxFirst}, {minSecond, maxSecond}})

	}
	partOne()
	partTwo()
}

func partOne() {
	total := 0
	for _, pair := range pairs {
		minFirst, maxFirst, minSecond, maxSecond := pair[0][0], pair[0][1], pair[1][0], pair[1][1]
		if (minFirst <= minSecond && maxSecond <= maxFirst) || (minSecond <= minFirst && maxFirst <= maxSecond) {
			total++
		}
	}
	fmt.Println(total)
}

func partTwo() {
	total := 0
	for _, pair := range pairs {
		seen := make(map[int64]bool)
		minFirst, maxFirst, minSecond, maxSecond := pair[0][0], pair[0][1], pair[1][0], pair[1][1]
		for i := minFirst; i <= maxFirst; i++ {
			seen[i] = true
		}
		for i := minSecond; i <= maxSecond; i++ {
			if seen[i] {
				total++
				break
			}
		}
	}
	fmt.Println(total)
}

func minMaxPair(pair string) (int64, int64) {
	var split = strings.Split(pair, "-")
	min, err := strconv.ParseInt(split[0], 10, 64)
	if err != nil {
		panic(err)
	}
	max, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		panic(err)
	}
	return min, max
}
