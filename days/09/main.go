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
	histories := parseHistories(lines)

	nextValues := make([]int, len(histories))
	for i, history := range histories {
		nextValues[i] = extrapolateNextValue(history, Forward)
	}

	return sumOfIntSlice(nextValues)
}

func part2(lines []string) int {
	histories := parseHistories(lines)
	previousValues := make([]int, len(histories))

	for i, history := range histories {
		previousValues[i] = extrapolateNextValue(history, Backward)
	}

	return sumOfIntSlice(previousValues)
}

func parseHistories(lines []string) [][]int {
	histories := make([][]int, len(lines))

	for i, line := range lines {
		s := strings.Fields(line)
		history := make([]int, len(s))

		for j, v := range s {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

			history[j] = num
		}

		histories[i] = history
	}

	return histories
}

const (
	Forward = iota
	Backward
)

func extrapolateNextValue(history []int, direction int) int {
	var values []int

	save := getValueInSequence(direction, history)
	values = append(values, save)
	sequence := history

	for !isAllZeroes(sequence) {
		sequence = getSequenceOfDifferences(sequence)
		save = getValueInSequence(direction, sequence)
		values = append(values, save)
	}

	next := 0
	for i := len(values) - 1; i > 0; i-- {
		if direction == Backward {
			next *= -1
		}

		next = values[i-1] + next
	}

	return next
}

func getValueInSequence(direction int, sequence []int) int {
	var value int

	switch direction {
	case Forward:
		// If extrapolating forward, return the value at the end of sequence.
		value = sequence[len(sequence)-1]
	case Backward:
		// If extrapolating backwards, return the value at the beginning of sequence.
		value = sequence[0]
	default:
		// Raise error if it's not either direction.
		log.Fatal(fmt.Sprintf("Error: Cannot read direction \"%v\"\n", direction))
	}

	return value
}

func getSequenceOfDifferences(sequence []int) []int {
	sequenceOfDifferences := make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		sequenceOfDifferences[i] = sequence[i+1] - sequence[i]
	}

	return sequenceOfDifferences
}

func isAllZeroes(intSlice []int) bool {
	for _, v := range intSlice {
		if v != 0 {
			return false
		}
	}

	return true
}

func sumOfIntSlice(intSlice []int) int {
	sum := 0
	for _, v := range intSlice {
		sum += v
	}
	return sum
}
