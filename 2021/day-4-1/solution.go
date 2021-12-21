package main

import (
	"bufio"
	"log"
	"regexp"
	"strings"

	"github.com/valderrama/advent-of-code/util/convert"
	"github.com/valderrama/advent-of-code/util/file"
)

const (
	inputFile = "input.txt"
	numRows   = 5
	numCols   = 5
)

var (
	splitRegexp = regexp.MustCompile(`\s+`)
)

type boardState struct {
	currValue int
	rowCounts [numRows]int
	colCounts [numCols]int
}

type boardEntry struct {
	boardID int
	row     int
	col     int
}

func main() {
	// Variables
	numberToBoardEntry := map[int][]*boardEntry{}
	boards := []*boardState{}
	numbers := []int{}

	// Open file
	file := file.Open(inputFile)
	defer file.Close()

	// For each line
	var currBoard *boardState
	boardID := -1
	row := 0
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		// First line we need to get the numbers
		if i == 0 {
			numberStrs := strings.Split(line, ",")
			for _, str := range numberStrs {
				numbers = append(numbers, convert.MustAtoi(str))
			}
			continue
		}

		switch (i - 1) % (numRows + 1) {

		// Empty line means new board
		case 0:
			if line != "" {
				log.Fatalf("line number %d should be empty instead of: %s", i, line)
			}
			// Reset board vars and add a board
			boardID++
			row = 0
			currBoard = &boardState{}
			boards = append(boards, currBoard)

		// Any other line should be parsed as a row
		default:
			// Ensure we have right number of rows
			if row >= numRows {
				log.Fatalf("line %d, board %d, row %d - only %d rows expected", i, boardID, row, numRows)
			}

			// Split the row by whitespace
			currRowNums := splitRegexp.Split(strings.Trim(line, " "), -1)

			// Ensure we have right number of columns
			if len(currRowNums) != numCols {
				log.Fatalf("line %d, board %d, row %d - %d cols, but only %d cols expected", i, boardID, row, len(currRowNums), numCols)
			}

			// For each column we get the number and create an entry then increment the board value
			for col, numStr := range currRowNums {
				num := convert.MustAtoi(numStr)
				numberToBoardEntry[num] = append(numberToBoardEntry[num], &boardEntry{boardID: boardID, row: row, col: col})
				currBoard.currValue += num
			}
			row++

		}
	}

	// Verify no I/O error
	if err := scanner.Err(); err != nil {
		log.Fatal("error while scanning: ", err)
	}

	// For each number
	for _, num := range numbers {
		// Get entries
		entries := numberToBoardEntry[num]
		// For each entry
		for _, entry := range entries {
			// Reduce the board value
			currBoard = boards[entry.boardID]
			currBoard.currValue -= num
			currBoard.colCounts[entry.col]++
			currBoard.rowCounts[entry.row]++
			if currBoard.colCounts[entry.col] >= numCols || currBoard.rowCounts[entry.row] >= numRows {
				log.Printf("Winner is board %d with a final score of %d * %d = %d", entry.boardID, currBoard.currValue, num, currBoard.currValue*num)
				return
			}
		}
	}

}
