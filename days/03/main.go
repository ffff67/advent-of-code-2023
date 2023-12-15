package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
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
	schematic := strings.Split(trimmedData, "\n")

	fmt.Printf("part 1: %d\n", part1(schematic))
	fmt.Printf("part 2: %d\n", part2(schematic))
}

type Coordinate struct {
	Row int
	Col int
}

func (c Coordinate) Equal(other Coordinate) bool {
	return c.Row == other.Row && c.Col == other.Col
}

func part1(schematic []string) int {
	sum := 0
	for i, row := range schematic {
		for j, r := range row {
			if isSymbol(r) {
				firstDigitIndices := checkSurroundings(schematic, Coordinate{Row: i, Col: j})
				numbers := getNumbers(schematic, firstDigitIndices)
				for _, n := range numbers {
					sum += n
				}
			}
		}
	}

	return sum
}

func part2(schematic []string) int {
	sum := 0
	for i, row := range schematic {
		for j, r := range row {
			if isSymbol(r) {
				firstDigitIndices := checkSurroundings(schematic, Coordinate{Row: i, Col: j})

				if len(firstDigitIndices) == 2 {
					numbers := getNumbers(schematic, firstDigitIndices)

					if len(numbers) == 2 {
						gearRatio := numbers[0] * numbers[1]
						sum += gearRatio
					}
				}
			}
		}
	}

	return sum
}

func isSymbol(x rune) bool {
	symbols := []rune{'@', '#', '$', '%', '&', '*', '+', '-', '/', '='}
	for _, sym := range symbols {
		if x == sym {
			return true
		}
	}

	return false
}

func checkSurroundings(schematic []string, coord Coordinate) []Coordinate {
	var indices []Coordinate
	length, width := len(schematic), len(schematic[0])

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			x := coord.Col + dx
			y := coord.Row + dy
			if x >= 0 && x <= width && y >= 0 && y <= length {
				if _, err := strconv.Atoi(string(schematic[y][x])); err == nil {
					for x > 0 && isDigit(schematic[y][x-1]) {
						x--
					}

					newIndex := Coordinate{Col: x, Row: y}
					if !slices.Contains(indices, newIndex) {
						indices = append(indices, newIndex)
					}
				}
			}
		}
	}

	return indices
}

func getNumbers(schematic []string, numberIndices []Coordinate) []int {
	var numbers []int
	var firstIndices []Coordinate

	for _, c := range numberIndices {
		var num, index int
		placeValue := 0
		for index = c.Col; index < len(schematic[c.Row]); index++ {
			if n, err := strconv.Atoi(string(schematic[c.Row][index])); err == nil {
				num *= 10
				num += n
				placeValue++
			} else {
				break
			}
		}
		firstIndex := Coordinate{Row: c.Row, Col: index}
		if !slices.Contains(firstIndices, firstIndex) {
			numbers = append(numbers, num)
			firstIndices = append(firstIndices, firstIndex)
		}
	}

	return numbers
}

func isDigit(b byte) bool {
	if _, err := strconv.Atoi(string(b)); err == nil {
		return true
	}

	return false
}
