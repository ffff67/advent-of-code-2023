package main

import (
	"bytes"
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

	lines := bytes.Split(data, []byte("\n"))

	fmt.Printf("part 1: %d\n", part1(lines))
	fmt.Printf("part 2: %d\n", part2(lines))
}

func part1(lines [][]byte) int {
	var calibrationValues []int
	for _, line := range lines {
		var values []int
		for _, b := range line {
			if num, err := strconv.Atoi(string(b)); err == nil {
				values = append(values, num)
			}
		}

		finalValue := values[0]*10 + values[len(values)-1]
		calibrationValues = append(calibrationValues, finalValue)
	}

	return getSum(calibrationValues)
}

func part2(lines [][]byte) int {
	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var calibrationValues []int
	for _, line := range lines {
		var values []int
		var substringIndexes []int
		for j, byte := range line {
			substring := string(line[:j+1])

			// Append number in word form to array of values.
			for k, v := range numberMap {
				if index := strings.LastIndex(substring, k); index != -1 {
					var isFoundBefore bool

					for _, si := range substringIndexes {
						if index == si {
							isFoundBefore = true
							break
						}
					}

					if !isFoundBefore {
						values = append(values, v)
						substringIndexes = append(substringIndexes, index)
					}
				}
			}

			// Append numbers in digit form to array of values.
			if n, err := strconv.Atoi(string(byte)); err == nil {
				values = append(values, n)
			}
		}

		finalValue := values[0]*10 + values[len(values)-1]
		calibrationValues = append(calibrationValues, finalValue)
	}

	return getSum(calibrationValues)
}

func getSum(values []int) int {
	var sum int
	for _, v := range values {
		sum += v
	}
	return sum
}
