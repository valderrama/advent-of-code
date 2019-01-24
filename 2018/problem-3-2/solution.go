package main

import (
	"fmt"
	"math"
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
type SquareInch uint16

const (
	// Free means the square inch of fabric has not been claimed
	Free SquareInch = 0
	// Conflict means the square inch of fabric has been claimed multiple times
	Conflict SquareInch = math.MaxUint16
)

// Claim represents an elf's claim over the fabric (i.e. one line in the input)
type Claim struct {
	ID     SquareInch
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
	potentialValidClaims := map[SquareInch]bool{}

	file.ProcessEachLine("input.txt", func(line string) bool {
		claim := parseClaim(re, line)
		potentialValidClaims[claim.ID] = true
		for col := claim.Col; col < claim.Col+claim.Width; col++ {
			for row := claim.Row; row < claim.Row+claim.Height; row++ {
				currSquareInch := &fabric[col][row]
				switch *currSquareInch {
				case Free:
					*currSquareInch = SquareInch(claim.ID)
				case Conflict:
					delete(potentialValidClaims, claim.ID)
				default:
					delete(potentialValidClaims, claim.ID)
					delete(potentialValidClaims, *currSquareInch)
					*currSquareInch = Conflict
				}
			}
		}
		return true
	})

	for key := range potentialValidClaims {
		fmt.Println(key)
	}
}

func parseClaim(re *regexp.Regexp, line string) Claim {
	match := re.FindStringSubmatch(line)
	return Claim{
		ID:     SquareInch(convert.MustParseUint16(match[1])), // Could verify is not 0 or max
		Col:    convert.MustAtoi(match[2]),
		Row:    convert.MustAtoi(match[3]),
		Width:  convert.MustAtoi(match[4]),
		Height: convert.MustAtoi(match[5]),
	}
}
