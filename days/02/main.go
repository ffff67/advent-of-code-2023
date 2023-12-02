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

	games := strings.Split(string(data), "\n")

	fmt.Printf("part 1: %d\n", part1(games))
	fmt.Printf("part 2: %d\n", part2(games))
}

func part1(games []string) int {
	const maxRed, maxGreen, maxBlue = 12, 13, 14

	var IDs []int

	for index, game := range games {
		sets := strings.Split(game[strings.Index(game, ": ")+2:], "; ")

		var isImpossibleBag bool

		for _, set := range sets {
			parsedSet := parseSet(set)
			if parsedSet["red"] > maxRed || parsedSet["green"] > maxGreen || parsedSet["blue"] > maxBlue {
				isImpossibleBag = true
				break
			}
		}

		if !isImpossibleBag {
			IDs = append(IDs, index+1)
		}
	}

	var sumIDs int
	for _, v := range IDs {
		sumIDs += v
	}
	return sumIDs
}

func part2(games []string) int {
	var powers []int

	for _, game := range games {
		sets := strings.Split(game[strings.Index(game, ": ")+2:], "; ")

		gameColorMap := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		for _, set := range sets {
			parsedSet := parseSet(set)
			for k, v := range parsedSet {
				if parsedSet[k] > gameColorMap[k] {
					gameColorMap[k] = v
				}
			}
		}

		power := gameColorMap["red"] * gameColorMap["blue"] * gameColorMap["green"]
		powers = append(powers, power)
	}

	var sum int
	for _, v := range powers {
		sum += v
	}
	return sum
}

func parseSet(set string) map[string]int {
	setMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	colors := strings.Split(set, ", ")
	for _, color := range colors {
		c := strings.Split(color, " ")
		n, err := strconv.Atoi(c[0])
		if err != nil {
			fmt.Println(err)
		}
		setMap[c[1]] = n
	}
	return setMap
}
