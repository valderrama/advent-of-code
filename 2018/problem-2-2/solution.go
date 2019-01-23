package main

import (
	"fmt"

	"github.com/valderrama/advent-of-code/util/file"
)

func main() {
	answer := ""
	substrMap := map[string]bool{}
	file.ProcessEachLine("input.txt", func(line string) bool {
		for index := range line {
			candidate := line[:index] + "_" + line[index+1:]
			if substrMap[candidate] {
				answer = line[:index] + line[index+1:]
				return false
			}
			substrMap[candidate] = true
		}
		return true
	})

	fmt.Println(answer)
}
