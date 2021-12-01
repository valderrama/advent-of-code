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
	file := file.Open(inputFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		log.Print(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error while scanning:", err)
	}
}
