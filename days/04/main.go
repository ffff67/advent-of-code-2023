package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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

	cards := strings.Split(string(data), "\n")

	fmt.Printf("part 1: %d\n", part1(cards))
}

func part1(cards []string) int {
	var totalPoints int

	re := regexp.MustCompile(`Card\s+\d+:\s+((?:\d+\s*)+)\s+\|\s+((?:\d+\s*)+)`)

	for _, card := range cards {
		matches := re.FindStringSubmatch(card)
		winningNumbers, yourNumbers := strings.Fields(matches[1]), strings.Fields(matches[2])

		matchingNumbers := matchingNumbers(convertToIntSlice(winningNumbers), convertToIntSlice(yourNumbers))

		points := calculatePoints(len(matchingNumbers))
		totalPoints += points
	}
	return totalPoints
}

func convertToIntSlice(s []string) []int {
	var res []int
	for _, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, n)
	}
	return res
}

func matchingNumbers(s1 []int, s2 []int) []int {
	var matches []int
	for _, n1 := range s1 {
		for _, n2 := range s2 {
			if n1 == n2 {
				matches = append(matches, n1)
			}
		}
	}
	return matches
}

func calculatePoints(matches int) int {
	if matches == 0 {
		return 0
	}

	points := 1
	for i := 0; i < matches-1; i++ {
		points *= 2
	}
	return points
}
