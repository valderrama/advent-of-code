package main

import (
	"bufio"
	"log"
	"regexp"

	"github.com/pkg/errors"
	"github.com/valderrama/advent-of-code/util/convert"
	"github.com/valderrama/advent-of-code/util/file"
)

const (
	inputFile = "input.txt"
	maxX      = 999
	maxY      = 999
)

var (
	lineRegexp = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func (l *line) points() []point {
	var xStep int
	var yStep int

	switch {
	case l.start.x < l.end.x:
		xStep = 1
	case l.start.x > l.end.x:
		xStep = -1
	default:
		xStep = 0
	}

	switch {
	case l.start.y < l.end.y:
		yStep = 1
	case l.start.y > l.end.y:
		yStep = -1
	default:
		yStep = 0
	}

	curr := l.start
	points := []point{curr}
	for curr != l.end {
		curr.x += xStep
		curr.y += yStep
		points = append(points, curr)
	}

	return points
}

func parseLine(text string) line {
	matches := lineRegexp.FindStringSubmatch(text)
	if matches == nil {
		panic(errors.Errorf("failed to match line regexp: %s", text))
	}

	if len(matches) != 5 {
		panic(errors.Errorf("expected match to have 4 submatches but it had: %d", len(matches)-1))
	}

	return line{
		start: point{
			x: convert.MustAtoi(matches[1]),
			y: convert.MustAtoi(matches[2]),
		},
		end: point{
			x: convert.MustAtoi(matches[3]),
			y: convert.MustAtoi(matches[4]),
		},
	}
}

type grid struct {
	impl         [maxX][maxY]int
	dangerPoints int
}

func (g *grid) applyLine(l line) {
	points := l.points()
	for _, point := range points {
		g.impl[point.x][point.y]++
		if g.impl[point.x][point.y] == 2 {
			g.dangerPoints++
		}
	}
}

func main() {
	// Variables
	var g grid

	// Open file
	file := file.Open(inputFile)
	defer file.Close()

	// For each line
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		text := scanner.Text()
		l := parseLine(text)
		g.applyLine(l)
	}

	// Verify no I/O error
	if err := scanner.Err(); err != nil {
		log.Fatal("error while scanning: ", err)
	}

	// Print answer
	log.Println(g.dangerPoints)
}
