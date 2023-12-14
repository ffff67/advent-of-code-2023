package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type RockType int

const (
	RoundRock RockType = iota
	CubeRock
	NoRock
)

var rockMap = map[rune]RockType{
	'O': RoundRock,
	'#': CubeRock,
	'.': NoRock,
}

func (r RockType) String() string {
	switch r {
	case RoundRock:
		return "⚫"
	case CubeRock:
		return "🟥"
	case NoRock:
		return "⬜"
	default:
		return "❓"
	}
}

type Platform [][]RockType

func (p Platform) String() string {
	var strs []string
	for i, row := range p {
		for _, r := range row {
			strs = append(strs, fmt.Sprintf("%s", r.String()))
		}
		strs = append(strs, fmt.Sprintf(" %d\n", len(p)-i))
	}

	return strings.Join(strs, "")
}

func (p Platform) Equal(other Platform) bool {
	for i, row := range p {
		for j, r := range row {
			if r != other[i][j] {
				return false
			}
		}
	}

	return true
}

type Direction int

const (
	North Direction = iota
	West
	South
	East
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
	platform := parseData(lines)
	shiftedPlatform := shiftPlatform(platform, North)

	return getTotalLoad(shiftedPlatform)
}

func part2(lines []string) int {
	platform := parseData(lines)

	var seen []Platform
	seen = appendPlatform(seen, platform)

	totalCycles := 1_000_000_000
	cycles := 0
	for {
		platform = shiftPlatform(platform, North)
		platform = shiftPlatform(platform, West)
		platform = shiftPlatform(platform, South)
		platform = shiftPlatform(platform, East)
		cycles++

		if containsPlatform(seen, platform) {
			break
		} else {
			seen = appendPlatform(seen, platform)
		}
	}

	offset := getIndexOfPlatform(seen, platform)
	loopLength := cycles - offset
	platform = seen[(totalCycles-offset)%loopLength+offset]

	return getTotalLoad(platform)
}

func parseData(lines []string) Platform {
	platform := make(Platform, len(lines))
	for i, line := range lines {
		platformRow := make([]RockType, len(line))
		for j, v := range line {
			rockType, ok := rockMap[v]
			if !ok {
				log.Fatal(fmt.Sprintf("Cannot read rock: %v", v))
			}
			platformRow[j] = rockType
		}
		platform[i] = platformRow
	}

	return platform
}

func shiftPlatform(p Platform, d Direction) Platform {

	if d == North || d == West {
		for rowN, row := range p {
			for colN, r := range row {
				if r == RoundRock {
					p = moveRock(p, rowN, colN, d)
				}
			}

		}
	} else if d == South || d == East {
		for rowN := len(p) - 1; rowN >= 0; rowN-- {
			for colN := len(p[0]) - 1; colN >= 0; colN-- {
				if p[rowN][colN] == RoundRock {
					p = moveRock(p, rowN, colN, d)
				}

			}
		}
	}

	return p
}

func moveRock(p Platform, rockRow, rockCol int, d Direction) Platform {
	p[rockRow][rockCol] = NoRock
	move := 0
	switch d {
	case North:
		for i := rockRow - 1; i >= 0; i-- {
			if p[i][rockCol] == NoRock {
				move++
			} else {
				break
			}
		}
		p[rockRow-move][rockCol] = RoundRock
	case West:
		for i := rockCol - 1; i >= 0; i-- {
			if p[rockRow][i] == NoRock {
				move++
			} else {
				break
			}
		}
		p[rockRow][rockCol-move] = RoundRock
	case South:
		for i := rockRow + 1; i < len(p); i++ {
			if p[i][rockCol] == NoRock {
				move++
			} else {
				break
			}
		}
		p[rockRow+move][rockCol] = RoundRock
	case East:
		for i := rockCol + 1; i < len(p[rockRow]); i++ {
			if p[rockRow][i] == NoRock {
				move++
			} else {
				break
			}
		}
		p[rockRow][rockCol+move] = RoundRock
	default:
		log.Fatal("Error: Cannot read direction")
	}

	return p
}

func getTotalLoad(p Platform) int {
	sum := 0
	for i, row := range p {
		for _, r := range row {
			if r == RoundRock {
				load := len(p) - i
				sum += load
			}
		}
	}

	return sum
}

func containsPlatform(ps []Platform, p Platform) bool {
	for _, pp := range ps {
		if p.Equal(pp) {
			return true
		}
	}

	return false
}

func getIndexOfPlatform(ps []Platform, p Platform) int {
	for i, pp := range ps {
		if p.Equal(pp) {
			return i
		}
	}
	return -1
}

func appendPlatform(ps []Platform, platform Platform) []Platform {
	p := make([][]RockType, len(platform))
	for i := range platform {
		p[i] = make([]RockType, len(platform[i]))
		copy(p[i], platform[i])
	}
	ps = append(ps, p)
	return ps
}
