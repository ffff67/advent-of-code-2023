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
		nextValues[i] = getNextValue(history)
	}

	return sumOfIntSlice(nextValues)
}

func part2(lines []string) int {
	histories := parseHistories(lines)

	previousValues := make([]int, len(histories))
	for i, history := range histories {
		previousValues[i] = getPreviousValue(history)
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

func getNextValue(history []int) int {
	var lastValues []int
	lastValues = append(lastValues, history[len(history)-1])

	sequence := history
	for !isAllZeroes(sequence) {
		sequence = getSequenceOfDifferences(sequence)
		lastValues = append(lastValues, sequence[len(sequence)-1])
	}

	next := 0
	for i := len(lastValues) - 1; i > 0; i-- {
		next = lastValues[i-1] + next
	}

	return next
}

func getPreviousValue(history []int) int {
	var firstValues []int
	firstValues = append(firstValues, history[0])

	sequence := history
	for !isAllZeroes(sequence) {
		sequence = getSequenceOfDifferences(sequence)
		firstValues = append(firstValues, sequence[0])
	}

	next := 0
	for i := len(firstValues) - 1; i > 0; i-- {
		next = firstValues[i-1] - next
	}

	return next
}

func getSequenceOfDifferences(history []int) []int {
	sequenceOfDifferences := make([]int, len(history)-1)
	for i := 0; i < len(history)-1; i++ {
		sequenceOfDifferences[i] = history[i+1] - history[i]
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
