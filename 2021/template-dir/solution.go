package main

import (
	"bufio"
	"log"

	"github.com/valderrama/advent-of-code/util/file"
)

const (
	inputFile = "input.txt"
)

func main() {
	// Variables

	// Open file
	file := file.Open(inputFile)
	defer file.Close()

	// For each line
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		log.Print(line)
	}

	// Verify no I/O error
	if err := scanner.Err(); err != nil {
		log.Fatal("error while scanning: ", err)
	}
}
