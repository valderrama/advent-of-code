package file

import (
	"bufio"
	"log"
	"os"
)

// Open will always return a file or exit, don't forget to close() it!
func Open(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

// ProcessEachLine will open a file and use the lineProcessor function to process each line
// to terminate early lineProcessor may return false
func ProcessEachLine(filename string, lineProcessor func(string) bool) {
	file := Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		shouldContinue := lineProcessor(scanner.Text())
		if !shouldContinue {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
