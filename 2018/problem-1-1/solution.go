package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/valderrama/advent-of-code/util/file"
)

func main() {
	total := 0
	file.ProcessEachLine("input.txt", func(line string) bool {
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		total += value
		return true
	})
	fmt.Println(total)
}
