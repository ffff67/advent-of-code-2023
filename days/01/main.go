package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input")
	data, _ := io.ReadAll(file)
	lines := bytes.Split(data, []byte("\n"))
	part1(lines)
	part2(lines)
}

func part1(lines [][]byte) {
	var calibrationValues []int
	for _, line := range lines {
		var value []string
		for j, byte := range line {
			s := string(byte)

			if _, err := strconv.Atoi(s); err == nil {
				value = append(value, s)
			}

			if j == (len(line) - 1) {
				recoveredValue, _ := strconv.Atoi(value[0] + value[len(value)-1])
				calibrationValues = append(calibrationValues, recoveredValue)
			}
		}
	}
	var sum int
	for _, v := range calibrationValues {
		sum += v
	}
	fmt.Printf("part 1: %d\n", sum)
}

func part2(lines [][]byte) {
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
		var value []int
		var substring string
		var substringIndexes []int
		for _, byte := range line {
			s := string(byte)

			substring += s
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
						value = append(value, v)
						substringIndexes = append(substringIndexes, index)
					}
				}
			}

			if n, err := strconv.Atoi(s); err == nil {
				value = append(value, n)
			}
		}

		finalValue, _ := strconv.Atoi(fmt.Sprintf("%v%v", value[0], value[len(value)-1]))

		calibrationValues = append(calibrationValues, finalValue)
	}

	var sum int
	for _, v := range calibrationValues {
		sum += v
	}
	fmt.Printf("part 2: %d\n", sum)
}
