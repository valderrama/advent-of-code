package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/valderrama/advent-of-code/util/file"
)

func main() {
	total := 0
	duplicate := 0
	hasDuplicate := false
	totals := make(map[int]bool)

	for !hasDuplicate {
		file.ProcessEachLine("input.txt", func(line string) bool {
			value, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			total += value

			if totals[total] {
				duplicate = total
				hasDuplicate = true
				return false
			}

			totals[total] = true
			return true
		})
	}

	fmt.Println(duplicate)
}
