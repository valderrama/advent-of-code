package main

import (
	"fmt"

	"github.com/valderrama/advent-of-code/util/file"
)

func main() {
	fmt.Println("Hi!")

	file.ProcessEachLine("input.txt", func(line string) bool {
		return true
	})
}
