package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
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

	lines := strings.Split(string(data), "\n")

	fmt.Printf("part 1: %d\n", part1(lines))
	fmt.Printf("part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	var locations []int
	seeds := parseSeeds(lines)
	maps := parseData(lines)

	for _, seed := range seeds {
		new := seed
		for _, m := range maps {
			new = sourceToDestination(m, new)
		}
		locations = append(locations, new)
	}

	return slices.Min(locations)
}

func part2(lines []string) int {
	seeds := parseSeedsPart2(lines)
	maps := parseData(lines)

	for _, ms := range maps {
		var new [][]int

		for len(seeds) > 0 {
			// Pop out seed
			seed := seeds[len(seeds)-1]
			seeds = seeds[:len(seeds)-1]

			seedStart, seedEnd := seed[0], seed[1]
			finishEarly := false

			for _, m := range ms {
				mStart, mEnd := m.Source, m.Source+m.Length
				overlapStart := max(seedStart, mStart)
				overlapEnd := min(seedEnd, mEnd)

				if overlapStart < overlapEnd {
					// If map range is within the seed range, append the converted values to new
					new = append(new, []int{overlapStart - m.Source + m.Dest, overlapEnd - m.Source + m.Dest})

					if overlapStart > seedStart {
						// The left end of the seed range that isn't in the overlap. Append it back to seeds and repeat the process of checking for overlaps.
						seeds = append(seeds, []int{seedStart, overlapStart})
					}
					if overlapEnd < seedEnd {
						// The right end of the seed range that isn't in the overlap. Append it back to seeds and repeat the process of checking for overlaps.
						seeds = append(seeds, []int{overlapEnd, seedEnd})
					}

					// If overlap is found, the loop should finish early
					finishEarly = true
					break
				}
			}

			if !finishEarly {
				// If there are no overlaps anymore, no conversion is needed and the seed can be appended as is.
				new = append(new, seed)
			}
		}

		seeds = new
	}

	minSeeds := make([]int, len(seeds))
	for i, s := range seeds {
		minSeeds[i] = s[0]
	}

	return slices.Min(minSeeds)
}

type SourceToDestMap struct {
	Source int
	Dest   int
	Length int
}

func sourceToDestination(s2d []SourceToDestMap, seed int) int {
	final := 0
	for _, stdm := range s2d {
		end := stdm.Source + stdm.Length - 1
		if seed >= stdm.Source && seed <= end {
			final = seed + (stdm.Dest - stdm.Source)
			break
		} else {
			final = seed
		}
	}

	return final
}

func parseSeeds(lines []string) []int {
	_, text, _ := strings.Cut(lines[0], " ")
	return convertToIntSlice(strings.Fields(text))
}

func parseSeedsPart2(lines []string) [][]int {
	_, text, _ := strings.Cut(lines[0], " ")
	nums := convertToIntSlice(strings.Fields(text))
	var res [][]int
	for index, n := range nums {
		if index%2 == 0 {
			res = append(res, []int{n, n + nums[index+1]})
		}
	}
	return res
}

func parseData(lines []string) [][]SourceToDestMap {
	var sourceToDests [][]SourceToDestMap
	var begin bool
	var pogMap []SourceToDestMap
	for _, line := range lines[2:] {
		if line == "" {
			begin = false
			sourceToDests = append(sourceToDests, pogMap)
			pogMap = []SourceToDestMap{}
		}

		if begin {
			std := strings.Split(line, " ")
			src, _ := strconv.Atoi(std[1])
			dest, _ := strconv.Atoi(std[0])
			length, _ := strconv.Atoi(std[2])
			stdm := SourceToDestMap{
				Source: src,
				Dest:   dest,
				Length: length,
			}
			pogMap = append(pogMap, stdm)
		}

		if strings.Contains(line, "map") {
			begin = true
		}
	}

	return sourceToDests
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
