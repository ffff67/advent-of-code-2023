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
	fmt.Printf("part 2: %d\n", part2(cards))

}

func part1(cards []string) int {
	var totalPoints int

	matches := getNumberOfMatchesForEachCard(cards)
	for _, numOfMatches := range matches {
		totalPoints += calculatePoints(numOfMatches)
	}
	return totalPoints
}

func part2(cards []string) int {
	matches := getNumberOfMatchesForEachCard(cards)
	var sum int
	for i := 0; i < len(matches); i++ {
		sum += countCards(matches, i)
	}

	return sum
}

func countCards(matches []int, cardNumber int) int {
	// Get number of copies
	copies := matches[cardNumber]

	// Recursively add copies for each copy
	var sum int
	for i := 0; i < copies; i++ {
		sum += countCards(matches, cardNumber+i+1)
	}

	// Add current card
	return 1 + sum
}

func getNumberOfMatchesForEachCard(cards []string) []int {
	var numOfMatchesSlice []int
	re := regexp.MustCompile(`Card\s+\d+:\s+((?:\d+\s*)+)\s+\|\s+((?:\d+\s*)+)`)

	for _, card := range cards {
		matches := re.FindStringSubmatch(card)
		winningNumbers, yourNumbers := strings.Fields(matches[1]), strings.Fields(matches[2])

		n := numberOfMatches(convertToIntSlice(winningNumbers), convertToIntSlice(yourNumbers))

		numOfMatchesSlice = append(numOfMatchesSlice, n)
	}

	return numOfMatchesSlice
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

func numberOfMatches(s1 []int, s2 []int) int {
	var matches int
	for _, n1 := range s1 {
		for _, n2 := range s2 {
			if n1 == n2 {
				matches += 1
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
