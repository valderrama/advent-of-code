package main

import (
	"bufio"
	"log"
	"strconv"

	"github.com/valderrama/advent-of-code/util/file"
)

const (
	inputFile = "input.txt"
	numBits   = 12
)

func main() {
	// Variables
	zeroLines := []string{}
	oneLines := []string{}

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
		switch line[0] {
		case '0':
			zeroLines = append(zeroLines, line)
		case '1':
			oneLines = append(oneLines, line)
		}
	}

	// Verify no I/O error
	if err := scanner.Err(); err != nil {
		log.Fatal("error while scanning: ", err)
	}

	var o2Candidates, co2Candidates []string
	if len(zeroLines) > len(oneLines) {
		o2Candidates = zeroLines
		co2Candidates = oneLines
	} else {
		o2Candidates = oneLines
		co2Candidates = zeroLines
	}

	o2 := -1
	for bitPos := 1; bitPos < numBits; bitPos++ {
		zeroLines = []string{}
		oneLines = []string{}
		for _, candidate := range o2Candidates {
			switch candidate[bitPos] {
			case '0':
				zeroLines = append(zeroLines, candidate)
			case '1':
				oneLines = append(oneLines, candidate)
			}
		}

		if len(zeroLines) > len(oneLines) {
			o2Candidates = zeroLines
		} else {
			o2Candidates = oneLines
		}

		if len(o2Candidates) == 1 {
			o264, err := strconv.ParseInt(o2Candidates[0], 2, 0)
			if err != nil {
				panic(err)
			}
			o2 = int(o264)
			break
		}
	}
	if o2 < 0 {
		log.Fatal("Did not find O2 value")
	}

	co2 := -1
	for bitPos := 1; bitPos < numBits; bitPos++ {
		zeroLines = []string{}
		oneLines = []string{}
		for _, candidate := range co2Candidates {
			switch candidate[bitPos] {
			case '0':
				zeroLines = append(zeroLines, candidate)
			case '1':
				oneLines = append(oneLines, candidate)
			}
		}

		if len(oneLines) < len(zeroLines) {
			co2Candidates = oneLines
		} else {
			co2Candidates = zeroLines
		}

		if len(co2Candidates) == 1 {
			co264, err := strconv.ParseInt(co2Candidates[0], 2, 0)
			if err != nil {
				panic(err)
			}
			co2 = int(co264)
			break
		}
	}
	if co2 < 0 {
		log.Fatal("Did not find CO2 value")
	}

	log.Printf("Result is %d * %d = %d", o2, co2, o2*co2)
}
