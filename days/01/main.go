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

	part1(lines)
	part2(lines)
}

func part1(lines [][]byte) {
	var calibrationValues []int
	for _, line := range lines {
		var value []string
		for _, b := range line {
			s := string(b)

			if _, err := strconv.Atoi(s); err == nil {
				value = append(value, s)
			}
		}

		recoveredValue, _ := strconv.Atoi(value[0] + value[len(value)-1])
		calibrationValues = append(calibrationValues, recoveredValue)
	}

	sum := getSum(calibrationValues)
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

		finalValue, err := strconv.Atoi(fmt.Sprintf("%v%v", value[0], value[len(value)-1]))
		if err != nil {
			log.Fatal(err)
		}

		calibrationValues = append(calibrationValues, finalValue)
	}

	sum := getSum(calibrationValues)
	fmt.Printf("part 2: %d\n", sum)
}

func getSum(values []int) int {
	var sum int
	for _, v := range values {
		sum += v
	}
	return sum
}
