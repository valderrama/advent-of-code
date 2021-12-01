package main

import (
	"bufio"
	"log"

	"github.com/valderrama/advent-of-code/util/convert"
	"github.com/valderrama/advent-of-code/util/file"
)

const (
	inputFile  = "input.txt"
	windowSize = 3
)

func sumIntSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

func main() {
	count := 0
	previousSum := 0
	window := [windowSize]int{}

	file := file.Open(inputFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		current := convert.MustAtoi(line)

		window[i%windowSize] = current

		currentSum := sumIntSlice(window[:])

		if i >= windowSize && currentSum > previousSum {
			count++
		}

		previousSum = currentSum
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("error while scanning: ", err)
	}

	log.Print("count is: ", count)
}
