package main

import (
	"fmt"

	"github.com/valderrama/advent-of-code/util/file"
)

func main() {
	twoCount := 0
	threeCount := 0
	file.ProcessEachLine("input.txt", func(line string) bool {
		letterCounts := map[rune]int{}
		localTwoCount := 0
		localThreeCount := 0
		for _, c := range line {
			val, ok := letterCounts[c]
			if !ok {
				letterCounts[c] = 1
			} else {
				switch val {
				case 1:
					localTwoCount++
				case 2:
					localTwoCount--
					localThreeCount++
				case 3:
					localThreeCount--
				}
				letterCounts[c] = val + 1
			}
		}
		if localTwoCount != 0 {
			twoCount++
		}
		if localThreeCount != 0 {
			threeCount++
		}
		return true
	})

	fmt.Println(twoCount * threeCount)
}
