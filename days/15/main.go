package main

import (
	"fmt"
	"io"
	"log"
	"os"
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
	lines := strings.Split(trimmedData, "\n")

	fmt.Printf("part 1: %d\n", part1(lines))
	fmt.Printf("part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	sequence := parseData(lines)
	sum := 0
	for _, step := range sequence {
		sum += hashAlgorithm(step)
	}

	return sum
}

type Lens struct {
	Label       string
	FocalLength int
}

func part2(lines []string) int {
	sequence := parseData(lines)
	boxes := make([][]Lens, 256)
	for _, step := range sequence {
		before, after, found := strings.Cut(step, "=")
		if found {
			boxNum := hashAlgorithm(before)
			label := before
			focalLength, err := strconv.Atoi(after)
			if err != nil {
				log.Fatal(err)
			}
			newLens := Lens{Label: label, FocalLength: focalLength}
			boxes = addLens(boxes, boxNum, newLens)

		} else {
			before, _, _ := strings.Cut(step, "-")
			boxNum := hashAlgorithm(before)
			label := before

			// remove lens with the given label from the box
			boxes = removeLens(boxes, boxNum, label)
		}
	}
	return totalFocusingPower(boxes)
}

func parseData(lines []string) []string {
	return strings.Split(lines[0], ",")
}

func hashAlgorithm(step string) int {
	value := 0
	for _, char := range step {
		// Increase value by the ASCII code of the character.
		value += int(char)

		// Multiply value by 17
		value *= 17

		// Set value to remainder of dividing itself by 256.
		value = value % 256
	}

	return value
}

func removeLens(boxes [][]Lens, boxNum int, label string) [][]Lens {
	box := boxes[boxNum]
	labelIndex := -1
	for i, lens := range box {
		if lens.Label == label {
			labelIndex = i
			break
		}
	}

	if labelIndex >= 0 {
		box = append(box[:labelIndex], box[labelIndex+1:]...)
	}

	boxes[boxNum] = box
	return boxes
}

func addLens(boxes [][]Lens, boxNum int, newLens Lens) [][]Lens {
	box := boxes[boxNum]
	alreadyInBox := false
	for i, lens := range box {
		if lens.Label == newLens.Label {
			// If there is already a lens in the box with the
			// same label, replace the old lens with the new
			// one.
			alreadyInBox = true
			box[i] = newLens
		}
	}

	if !alreadyInBox {
		// If there is not already a lens in the box with the
		// same label, add the lens to the box immediately
		// behind any lenses already in the box.
		box = append(box, newLens)
		boxes[boxNum] = box
	}

	return boxes
}

func totalFocusingPower(boxes [][]Lens) int {
	sum := 0
	for i, box := range boxes {
		for j, lens := range box {
			power := (1 + i) * (j + 1) * lens.FocalLength
			sum += power
		}
	}

	return sum
}
