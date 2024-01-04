package main

import (
	"fmt"
	"io"
	"log"
	"os"
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
	_, galaxies, emptyRows, emptyCols := parseData(lines)
	galaxies = updateGalaxies(galaxies, emptyRows, emptyCols, 2)
	pairsOfGalaxies := getPairsOfGalaxies(len(galaxies))

	sum := 0
	for _, pair := range pairsOfGalaxies {
		sum += getLengthOfShortestPath(galaxies, pair)
	}

	return sum
}

func part2(lines []string) int {
	_, galaxies, emptyRows, emptyCols := parseData(lines)
	galaxies = updateGalaxies(galaxies, emptyRows, emptyCols, 1_000_000)
	pairsOfGalaxies := getPairsOfGalaxies(len(galaxies))

	sum := 0
	for _, pair := range pairsOfGalaxies {
		sum += getLengthOfShortestPath(galaxies, pair)
	}

	return sum
}

type Galaxy struct {
	Number int
	Row    int
	Col    int
}

func parseData(lines []string) ([][]int, []Galaxy, []int, []int) {
	image := make([][]int, len(lines))
	var emptyRows, emptyCols []int
	var galaxies []Galaxy
	galaxyCount := 0
	for i, line := range lines {
		row := make([]int, len(line))
		var rowNotEmpty bool
		for j, v := range line {
			if v == '#' {
				galaxyCount++
				galaxies = append(galaxies, Galaxy{
					Number: galaxyCount,
					Row:    i,
					Col:    j,
				})
				rowNotEmpty = true
			}
		}
		if !rowNotEmpty {
			emptyRows = append(emptyRows, i)
		}
		image[i] = row
	}

	// Get empty columns
	for i := 0; i < len(lines[0]); i++ {
		var colNotEmpty bool
		for j := 0; j < len(lines); j++ {
			if lines[j][i] != '.' {
				colNotEmpty = true
			}
		}
		if !colNotEmpty {
			emptyCols = append(emptyCols, i)
		}
	}

	return image, galaxies, emptyRows, emptyCols
}

func updateGalaxies(galaxies []Galaxy, emptyRows, emptyCols []int, expansionScale int) []Galaxy {
	for i, galaxy := range galaxies {
		for _, row := range emptyRows {
			if galaxy.Row > row {
				galaxies[i].Row += (expansionScale - 1)
			} else {
				break
			}
		}
		for _, col := range emptyCols {
			if galaxy.Col > col {
				galaxies[i].Col += (expansionScale - 1)
			} else {
				break
			}
		}
	}

	return galaxies
}

func getPairsOfGalaxies(numberOfGalaxies int) [][]int {
	var pairs [][]int
	for i := 1; i < numberOfGalaxies+1; i++ {
		for j := i + 1; j < numberOfGalaxies+1; j++ {
			pairs = append(pairs, []int{i, j})
		}
	}

	return pairs
}

func getLengthOfShortestPath(galaxies []Galaxy, pair []int) int {
	a := galaxies[pair[0]-1]
	b := galaxies[pair[1]-1]
	steps := max(a.Row-b.Row, b.Row-a.Row) + max(a.Col-b.Col, b.Col-a.Col)
	return steps
}
