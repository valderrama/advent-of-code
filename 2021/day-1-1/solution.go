package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/valderrama/advent-of-code/util/file"
)

func main() {
	var err error
	previous := math.MaxInt
	current := 0
	count := 0
	file.ProcessEachLine("input.txt", func(line string) bool {
		current, err = strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if current > previous {
			count++
		}

		previous = current
		return true
	})

	fmt.Printf("count is: %d", count)
	fmt.Println()
}
