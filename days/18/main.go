package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	U Direction = iota
	D
	L
	R
)

type Dig struct {
	Dir       Direction
	Distance  int
	HexaColor string
}

type Coordinate struct {
	Row int
	Col int
}

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
	fmt.Printf("part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	digPlan := parseLines(lines)
	edgePoints := followDigPlan(digPlan)

	// Shoelace formula
	area := getArea(edgePoints)

	// Pick's theorem
	b := 0
	for _, d := range digPlan {
		b += d.Distance
	}
	i := area - b/2 + 1

	return i + b
}

func part2(lines []string) int {
	digPlan := parseLines2(lines)
	edgePoints := followDigPlan(digPlan)

	// Shoelace formula
	area := getArea(edgePoints)

	// Pick's theorem
	b := 0
	for _, d := range digPlan {
		b += d.Distance
	}
	i := area - b/2 + 1

	return i + b
}

func parseLines(lines []string) []Dig {
	digs := make([]Dig, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		direction, err := getDirection(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		distance, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		color := fields[2][2:8]

		digs[i] = Dig{
			Dir:       direction,
			Distance:  distance,
			HexaColor: color,
		}
	}

	return digs
}

func parseLines2(lines []string) []Dig {
	digs := make([]Dig, len(lines))
	for i, line := range lines {
		hexa := strings.Fields(line)[2][2:8]
		direction, err := getDirection2(hexa[5])
		if err != nil {
			log.Fatal(err)
		}
		distance, err := decodeDistance(hexa[:5])
		if err != nil {
			log.Fatal(err)
		}

		digs[i] = Dig{
			Dir:      direction,
			Distance: distance,
		}
	}

	return digs
}

func getDirection(s string) (Direction, error) {
	switch s {
	case "U":
		return U, nil
	case "D":
		return D, nil
	case "L":
		return L, nil
	case "R":
		return R, nil
	default:
		return 0, fmt.Errorf("Cannot convert to Direction.")
	}
}

func getDirection2(b byte) (Direction, error) {
	m := map[byte]Direction{'0': R, '1': D, '2': L, '3': U}
	if d, found := m[b]; found {
		return d, nil
	}

	return 0, fmt.Errorf("Cannot convert to Direction.")
}

func decodeDistance(hexa string) (int, error) {
	if s, err := strconv.ParseInt(hexa, 16, 32); err == nil {
		return int(s), nil
	} else {
		return 0, err
	}
}

func followDigPlan(digPlan []Dig) []Coordinate {
	lagoonEdge := make([]Coordinate, len(digPlan)+1)
	currLocation := Coordinate{Row: 0, Col: 0}
	lagoonEdge[0] = currLocation

	for i, dig := range digPlan {
		row, col := currLocation.Row, currLocation.Col
		switch dig.Dir {
		case U:
			currLocation = Coordinate{row - dig.Distance, col}
		case D:
			currLocation = Coordinate{row + dig.Distance, col}
		case L:
			currLocation = Coordinate{row, col - dig.Distance}
		case R:
			currLocation = Coordinate{row, col + dig.Distance}
		}
		lagoonEdge[i+1] = currLocation
	}

	return lagoonEdge
}

func getArea(edgePoints []Coordinate) int {
	// Shoelace formula
	sum := 0
	length := len(edgePoints)
	for i, p := range edgePoints {
		sum += p.Row * (edgePoints[(i-1+length)%length].Col - edgePoints[(i+1)%length].Col)

	}
	area := int(math.Abs(float64(sum)) / 2)

	return area
}
