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
	hotSprings := parseData(lines)

	var sum int
	for _, springRow := range hotSprings {
		sum += getPossibleArrangements(springRow)
	}

	return sum
}

func part2(lines []string) int {
	hotSprings := parseData(lines)

	for i := range hotSprings {
		row, groups := hotSprings[i].RowStr, hotSprings[i].DamagedGroups

		rowSlices := make([]string, 5)
		for i := range rowSlices {
			rowSlices[i] = row
		}
		repeatedRow := strings.Join(rowSlices, "?")

		hotSprings[i] = SpringRow{
			RowStr:        repeatedRow,
			DamagedGroups: repeatIntSlice(groups, 5),
		}
	}

	var sum int
	for _, springRow := range hotSprings {
		sum += getPossibleArrangements(springRow)
	}

	return sum
}

type SpringRow struct {
	RowStr        string
	DamagedGroups []int
}

type SpringRowComparable struct {
	RowStr        string
	DamagedGroups string
}

func parseData(lines []string) []SpringRow {
	records := make([]SpringRow, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)

		var nums []int
		groups := strings.Split(fields[1], ",")
		for _, s := range groups {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, n)
		}

		records[i] = SpringRow{RowStr: fields[0], DamagedGroups: nums}
	}

	return records
}

func getPossibleArrangements(springRow SpringRow) int {
	row, groups := springRow.RowStr, springRow.DamagedGroups
	return count(row, groups)
}

var cache = map[SpringRowComparable]int{}

func count(rowStr string, groups []int) int {
	// Base case: If rowStr is empty and there are no more damaged groups,
	// then it is a possible arrangement. Therefore, add 1 to count.
	if rowStr == "" {
		if len(groups) == 0 {
			return 1
		} else {
			return 0
		}
	}

	// Base case: If there are no more damaged groups and the remaining row
	// doesn't have a damaged spring, then it is a possible arrangement.
	// Therefore, add 1 to count.
	if len(groups) == 0 {
		if containsRune(rowStr, '#') {
			return 0
		} else {
			return 1
		}
	}

	// Create a key for caching purposes.
	key := SpringRowComparable{
		RowStr:        rowStr,
		DamagedGroups: fmt.Sprint(groups),
	}
	// Check if key is in the cache and if it is, return the result.
	if v, ok := cache[key]; ok {
		return v
	}

	res := 0
	firstRune := []rune(rowStr)[0]

	// If the first rune is a dot (operational spring) or a question mark
	// (unknown spring), skip it and recursively call count on the rest of the string.
	if containsRune(".?", firstRune) {
		res += count(rowStr[1:], groups)
	}

	// If the first rune is a broken spring (#) or an unknown spring (?),
	// and the conditions for the group is met, recursively call count on the rest of the string.
	// The conditions for the group are:
	// 1. There are enough springs left to fit in group
	// AND 2. The first N springs are all broken
	// AND 3. we reached the end of rowStr OR the next spring after group in rowStr
	// is an operational spring (.) which will separate this group from the next group.
	if containsRune("#?", firstRune) {
		if groups[0] <= len(rowStr) &&
			!containsRune(rowStr[:groups[0]], '.') &&
			(groups[0] == len(rowStr) || rowStr[groups[0]] != '#') {
			if groups[0] >= len(rowStr) {
				// If group is greater than or equal to rowStr, there is
				// nothing left in rowStr after group, so recursively call
				// count with an empty string.
				res += count("", groups[1:])
			} else {
				// Recursively call count with the rest of rowStr after group.
				// Add +1 to group in the startIndex of rowStr as there
				// is an operational spring (.) between damaged groups.
				res += count(rowStr[groups[0]+1:], groups[1:])
			}
		}
	}

	// Store the result in the cache.
	cache[key] = res

	return res
}

func containsRune(str string, r rune) bool {
	for _, s := range str {
		if s == r {
			return true
		}
	}

	return false
}

func repeatIntSlice(slice []int, times int) []int {
	var newSlice []int
	for i := 0; i < times; i++ {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}

	}

	return newSlice
}
