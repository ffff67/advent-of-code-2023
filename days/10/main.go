package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	trimmedData := strings.Trim(string(data), "\n")
	lines := strings.Split(trimmedData, "\n")

	fmt.Printf("part 1: %d\n", part1(lines))
}

type Unit int

func (u Unit) String() string {
	res := ""
	switch u {
	case 0:
		res = "|"
	case 1:
		res = "-"
	case 2:
		res = "L"
	case 3:
		res = "J"
	case 4:
		res = "7"
	case 5:
		res = "F"
	case 6:
		res = "."
	case 7:
		res = "S"
	default:
		res = "E"
	}

	return res
}

type Field [][]Unit

func (f Field) String() string {
	var res string
	for _, row := range f {
		for _, r := range row {
			res += r.String()
		}
		res += "\n"
	}
	return res
}

const (
	VerticalPipe Unit = iota
	HorizontalPipe
	NEBend
	NWBend
	SWBend
	SEBend
	Ground
	Animal
	Empty
)

type Coordinate struct {
	Row    int
	Column int
}

type Pipe struct {
	Unit       Unit
	Coordinate Coordinate
}

func part1(lines []string) int {
	field, animal := parseLines(lines)
	pipes := getPipes(field, animal)
	return len(pipes) / 2
}

func parseLines(lines []string) (Field, Pipe) {
	field := make([][]Unit, len(lines))
	var animal Pipe
	m := map[rune]Unit{
		'|': VerticalPipe,
		'-': HorizontalPipe,
		'L': NEBend,
		'J': NWBend,
		'7': SWBend,
		'F': SEBend,
		'.': Ground,
		'S': Animal,
	}

	for i, line := range lines {
		field[i] = make([]Unit, len(line))
		for j, v := range line {
			if unit, found := m[v]; found {
				field[i][j] = unit
				if unit == Animal {
					animal = Pipe{
						Unit: Animal,
						Coordinate: Coordinate{
							Row:    i,
							Column: j,
						},
					}
				}
			} else {
				log.Fatal(fmt.Sprintf("Cannot read unit at [%d, %d]", j, i))
			}
		}

	}

	return field, animal
}

func getPipes(field [][]Unit, animal Pipe) []Pipe {
	seen := []Pipe{animal}
	queue := []Pipe{animal}

	for len(queue) > 0 {
		out := queue[0]
		queue = queue[1:]
		row, col := out.Coordinate.Row, out.Coordinate.Column

		if canGoTop(field, out, seen) {
			seen = append(seen,
				Pipe{
					Unit:       field[row-1][col],
					Coordinate: Coordinate{Row: row - 1, Column: col},
				})

			queue = append(queue,
				Pipe{
					Unit:       field[row-1][col],
					Coordinate: Coordinate{Row: row - 1, Column: col},
				})
		}

		if canGoBottom(field, out, seen) {
			seen = append(seen,
				Pipe{
					Unit:       field[row+1][col],
					Coordinate: Coordinate{Row: row + 1, Column: col},
				})

			queue = append(queue,
				Pipe{
					Unit:       field[row+1][col],
					Coordinate: Coordinate{Row: row + 1, Column: col},
				})
		}

		if canGoLeft(field, out, seen) {
			seen = append(seen,
				Pipe{
					Unit:       field[row][col-1],
					Coordinate: Coordinate{Row: row, Column: col - 1},
				})

			queue = append(queue,
				Pipe{
					Unit:       field[row][col-1],
					Coordinate: Coordinate{Row: row, Column: col - 1},
				})
		}

		if canGoRight(field, out, seen) {
			seen = append(seen,
				Pipe{
					Unit:       field[row][col+1],
					Coordinate: Coordinate{Row: row, Column: col + 1},
				})

			queue = append(queue,
				Pipe{
					Unit:       field[row][col+1],
					Coordinate: Coordinate{Row: row, Column: col + 1},
				})
		}
	}

	return seen
}

func canGoTop(field [][]Unit, pipe Pipe, seen []Pipe) bool {
	unit, row, col := pipe.Unit, pipe.Coordinate.Row, pipe.Coordinate.Column

	if row > 0 && containsUnit([]Unit{Animal, VerticalPipe, NWBend, NEBend}, unit) && containsUnit([]Unit{VerticalPipe, SWBend, SEBend}, field[row-1][col]) && !containsPipe(seen, Pipe{Coordinate: Coordinate{Row: row - 1, Column: col}}) {
		return true
	}

	return false
}

func canGoBottom(field [][]Unit, pipe Pipe, seen []Pipe) bool {
	unit, row, col := pipe.Unit, pipe.Coordinate.Row, pipe.Coordinate.Column

	if row < len(field)-1 && containsUnit([]Unit{Animal, VerticalPipe, SWBend, SEBend}, unit) && containsUnit([]Unit{VerticalPipe, NWBend, NEBend}, field[row+1][col]) && !containsPipe(seen, Pipe{Coordinate: Coordinate{Row: row + 1, Column: col}}) {
		return true
	}

	return false
}

func canGoLeft(field [][]Unit, pipe Pipe, seen []Pipe) bool {
	unit, row, col := pipe.Unit, pipe.Coordinate.Row, pipe.Coordinate.Column

	if col > 0 && containsUnit([]Unit{Animal, HorizontalPipe, NWBend, SWBend}, unit) && containsUnit([]Unit{HorizontalPipe, NEBend, SEBend}, field[row][col-1]) && !containsPipe(seen, Pipe{Coordinate: Coordinate{Row: row, Column: col - 1}}) {
		return true
	}

	return false
}

func canGoRight(field [][]Unit, pipe Pipe, seen []Pipe) bool {
	unit, row, col := pipe.Unit, pipe.Coordinate.Row, pipe.Coordinate.Column

	if col < len(field[row])-1 && containsUnit([]Unit{Animal, HorizontalPipe, NEBend, SEBend}, unit) && containsUnit([]Unit{HorizontalPipe, NWBend, SWBend}, field[row][col+1]) && !containsPipe(seen, Pipe{Coordinate: Coordinate{Row: row, Column: col + 1}}) {
		return true
	}

	return false
}

func containsUnit(s []Unit, u Unit) bool {
	for _, a := range s {
		if a == u {
			return true
		}
	}

	return false
}

func containsPipe(s []Pipe, p Pipe) bool {
	for _, a := range s {
		if a.Coordinate.Row == p.Coordinate.Row && a.Coordinate.Column == p.Coordinate.Column {
			return true
		}
	}

	return false
}
