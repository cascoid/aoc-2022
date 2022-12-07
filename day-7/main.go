package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

const baseDir = "/"

var cmdLines []string
var directories = make(map[string][]string)
var folderSizes = make(map[string]int)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	cmdLines = strings.Split(string(i), "\n")

	mapDirectories()
	getFolderSizes(baseDir)
	partOne()
	partTwo()
}

func partOne() {
	total := 0
	for _, folderSize := range folderSizes {
		if folderSize <= 100000 {
			total += folderSize
		}
	}
	fmt.Println(total)
}

func partTwo() {
	spaceTaken := 70000000 - folderSizes[baseDir]
	spaceRequired := 30000000 - spaceTaken
	amountToDelete := 30000000
	for _, folderSize := range folderSizes {
		if spaceRequired <= folderSize && amountToDelete > folderSize {
			amountToDelete = folderSize
		}
	}
	fmt.Println(amountToDelete)
}

func getFolderSizes(parent string) int {
	total := 0
	for _, directory := range directories[parent] { // map over items in directory
		if strings.HasPrefix(directory, "dir ") {
			// recurse into subfolder
			subfolder := strings.Split(directory, " ")
			subfolderPath := filepath.Join(parent, subfolder[1])
			total += getFolderSizes(subfolderPath)
		} else {
			// its a file, add it to the total
			splitFile := strings.Split(directory, " ")
			size, _ := strconv.Atoi(splitFile[0])
			total += size
		}
	}
	folderSizes[parent] = total
	return total
}

func mapDirectories() {
	parent := ""
	for _, cmd := range cmdLines {
		if cmd == "$ ls" { // skip over ls command
			continue
		} else if strings.HasPrefix(cmd, "$ cd ") { // means directory needs changing
			directoryName := cmd[5:]   // get the directory name
			if directoryName == ".." { // dir name wants you to go back a directory
				parent = filepath.Dir(parent) // move up to parent for next iteration
			} else {
				parent = filepath.Clean(filepath.Join(parent, directoryName)) // move down to child for next iteration
			}
		} else {
			directories[parent] = append(directories[parent], cmd) // cmd is a file, add it to the parent directory map
		}
	}
}
