package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"unicode"
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

	schematic := strings.Split(string(data), "\n")

	fmt.Printf("part 1: %d\n", part1(schematic))
}

type VisualPoint struct {
	Value int
	X     int
	Y     int
}

func part1(schematic []string) int {
	var partNumbers []int

	for y, line := range schematic {
		var numbers []VisualPoint
		for x, r := range line {

			// Go down one by one element and store consecutive numbers
			if unicode.IsDigit(r) {
				vp := VisualPoint{
					Value: int(r - '0'),
					X:     x,
					Y:     y,
				}
				numbers = append(numbers, vp)
			} else {
				// If full number, check surroundings for symbols
				if len(numbers) > 0 && isAdjacentToSymbol(schematic, numbers) {
					// Add to array of part numbers
					finalValue := getFinalValue(numbers)
					partNumbers = append(partNumbers, finalValue)
				}

				numbers = []VisualPoint{}
			}
		}
		// Check number at the end of line
		if len(numbers) > 0 && isAdjacentToSymbol(schematic, numbers) {
			finalValue := getFinalValue(numbers)
			partNumbers = append(partNumbers, finalValue)
		}
	}

	return sum(partNumbers)
}

func getFinalValue(numbers []VisualPoint) int {
	var values []int
	for _, vp := range numbers {
		values = append(values, vp.Value)
	}
	finalValue := combineNumbers(values)
	return finalValue
}

func combineNumbers(nums []int) int {
	var res int
	l := len(nums)
	for i, v := range nums {
		pow := math.Pow(10, float64(l-i-1))
		res += v * int(pow)
	}
	return res
}

type Coordinate struct {
	X int
	Y int
}

func isAdjacentToSymbol(schematic []string, numbers []VisualPoint) bool {
	for _, vp := range numbers {
		// get array of surroundings
		var surroundings []Coordinate
		var left, right, up, down, diagonalUpLeft, diagonalUpRight, diagonalDownLeft, diagonalDownRight Coordinate
		width, height := len(schematic[0]), len(schematic)

		if vp.X > 0 {
			left = Coordinate{X: vp.X - 1, Y: vp.Y}
			surroundings = append(surroundings, left)
		}
		if vp.X < width-2 {
			right = Coordinate{X: vp.X + 1, Y: vp.Y}
			surroundings = append(surroundings, right)
		}
		if vp.Y > 0 {
			up = Coordinate{X: vp.X, Y: vp.Y - 1}
			surroundings = append(surroundings, up)
		}

		if vp.Y < height-2 {
			down = Coordinate{X: vp.X, Y: vp.Y + 1}
			surroundings = append(surroundings, down)
		}

		if vp.X > 0 && vp.Y > 0 {
			diagonalUpLeft = Coordinate{X: vp.X - 1, Y: vp.Y - 1}
			surroundings = append(surroundings, diagonalUpLeft)
		}

		if vp.X > 0 && vp.Y < height-2 {
			diagonalDownLeft = Coordinate{X: vp.X - 1, Y: vp.Y + 1}
			surroundings = append(surroundings, diagonalDownLeft)
		}

		if vp.X < width-2 && vp.Y > 0 {
			diagonalUpRight = Coordinate{X: vp.X + 1, Y: vp.Y - 1}
			surroundings = append(surroundings, diagonalUpRight)
		}

		if vp.X < width-2 && vp.Y < height-2 {
			diagonalDownRight = Coordinate{X: vp.X + 1, Y: vp.Y + 1}
			surroundings = append(surroundings, diagonalDownRight)
		}

		for _, coord := range surroundings {
			if isSymbol(string(schematic[coord.Y][coord.X])) {
				return true
			}
		}

	}

	return false
}

func isSymbol(x string) bool {
	symbols := []string{"@", "#", "$", "%", "&", "*", "+", "-", "/", "="}
	for _, sym := range symbols {
		if x == sym {
			return true
		}
	}

	return false
}

func sum(nums []int) int {
	var res int
	for _, v := range nums {
		res += v
	}
	return res
}
