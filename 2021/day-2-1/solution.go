package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/valderrama/advent-of-code/util/convert"
	"github.com/valderrama/advent-of-code/util/file"
)

const (
	inputFile = "input.txt"
)

const (
	up      = "up"
	down    = "down"
	forward = "forward"
)

func main() {
	x, y := 0, 0

	file := file.Open(inputFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			log.Fatalf("line split into %d parts instead of 2: %s", len(parts), parts)
		}

		direction, distance := parts[0], convert.MustAtoi(parts[1])
		switch direction {
		case up:
			y -= distance
		case down:
			y += distance
		case forward:
			x += distance
		default:
			log.Fatalf("invalid direction on line %d: %s", i+1, direction)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error while scanning: ", err)
	}

	log.Print("Result is: ", x*y)
}
