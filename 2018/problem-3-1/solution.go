package main

import (
	"fmt"
	"regexp"

	"github.com/valderrama/advent-of-code/util/convert"
	"github.com/valderrama/advent-of-code/util/file"
)

const (
	// FabricWidth is the width of the fabric for Santa's suit (i.e. # of columns)
	FabricWidth = 1000
	// FabricLength is the length of the fabric for Santa's suit (i.e. # of rows)
	FabricLength = 1000
)

// SquareInch is an enum used to describe the state of individual square inches of the fabric
type SquareInch uint8

const (
	// Free means the square inch of fabric has not been claimed
	Free SquareInch = iota
	// Claimed means the square inch of fabric has been claimed
	Claimed
	// Conflict means the square inch of fabric has been claimed multiple times
	Conflict
)

// Claim represents an elf's claim over the fabric (i.e. one line in the input)
type Claim struct {
	ID     string
	Col    int
	Row    int
	Width  int
	Height int
}

func main() {
	// access using [col][row]
	fabric := make([][]SquareInch, FabricWidth)
	for i := range fabric {
		fabric[i] = make([]SquareInch, FabricLength)
	}
	// Format: #ID @ COL,ROW: WIDTHxHEIGHT
	// Example: #1338 @ 983,969: 11x25
	re := regexp.MustCompile("#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	conflictTotal := 0

	file.ProcessEachLine("input.txt", func(line string) bool {
		claim := parseClaim(re, line)
		for col := claim.Col; col < claim.Col+claim.Width; col++ {
			for row := claim.Row; row < claim.Row+claim.Height; row++ {
				currSquareInch := &fabric[col][row]
				switch *currSquareInch {
				case Free:
					*currSquareInch = Claimed
				case Claimed:
					*currSquareInch = Conflict
					conflictTotal++
				case Conflict:
					// Nothing to do, already counted this square
				default:
					panic("Invalid SquareInch value: " + string(*currSquareInch))
				}
			}
		}
		return true
	})

	fmt.Println(conflictTotal)
}

func parseClaim(re *regexp.Regexp, line string) Claim {
	match := re.FindStringSubmatch(line)
	return Claim{
		ID:     match[1],
		Col:    convert.MustAtoi(match[2]),
		Row:    convert.MustAtoi(match[3]),
		Width:  convert.MustAtoi(match[4]),
		Height: convert.MustAtoi(match[5]),
	}
}
