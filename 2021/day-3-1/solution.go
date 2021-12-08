package main

import (
	"bufio"
	"log"

	"github.com/valderrama/advent-of-code/util/convert"
	"github.com/valderrama/advent-of-code/util/file"
)

const (
	inputFile = "input.txt"
	numBits   = 12
)

func main() {
	// Variables
	columnSums := [numBits]int{}

	// Open file
	file := file.Open(inputFile)
	defer file.Close()

	// For each line
	scanner := bufio.NewScanner(file)
	numLines := 0
	for ; scanner.Scan(); numLines++ {
		line := scanner.Text()

		// Should never happen
		if len(line) != numBits {
			log.Fatalf("line %d has %d bits instead of %d", numLines, len(line), numBits)
		}

		// Increment column sum
		for j := 0; j < numBits; j++ {
			columnSums[j] += convert.MustAtoi(string(line[j]))
		}
	}

	// Verify no I/O error
	if err := scanner.Err(); err != nil {
		log.Fatal("error while scanning: ", err)
	}

	g := 0
	e := 0
	for i := 0; i < numBits; i++ {
		// if our column sum is greater than half the number of lines then it is most common so
		// we need to set the g bit, otherwise we set the e bit
		if columnSums[i] > numLines/2 {
			// e.g. if numBits is 12 and i is 0 then we want to set the bit marked below
			// col  12 11 10 9  8  7  6  5  4  3  2  1
			// set  *
			// bit  0  0  0  0  0  0  0  0  0  0  0  0
			// this requires shifting numBits - 1 - 0 = 11 positions
			g |= (1 << (numBits - 1 - i))
		} else {
			e |= (1 << (numBits - 1 - i))
		}
	}

	log.Printf("Result is %d * %d = %d", g, e, g*e)
}
