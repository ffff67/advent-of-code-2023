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

	lines := strings.Split(string(data), "\n")

	fmt.Printf("part 1: %d\n", part1(lines))
	fmt.Printf("part 2: %d\n", part2(lines))
}

type Race struct {
	Time           int
	RecordDistance int
}

func part1(lines []string) int {
	time := convertToIntSlice(strings.Fields(lines[0])[1:])
	distance := convertToIntSlice(strings.Fields(lines[1])[1:])

	races := make([]Race, 0, len(time))
	for i := 0; i < len(time); i++ {
		races = append(races, Race{
			Time:           time[i],
			RecordDistance: distance[i],
		})
	}

	marginOfError := 1
	for _, race := range races {
		marginOfError *= waysToWin(race)
	}

	return marginOfError
}

func part2(lines []string) int {
	time := strings.ReplaceAll(lines[0][5:], " ", "")
	distance := strings.ReplaceAll(lines[1][9:], " ", "")

	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(distance)
	race := Race{
		Time:           t,
		RecordDistance: d,
	}

	return waysToWin2(race)
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

func waysToWin(race Race) int {
	var wins int
	for i := 0; i < race.Time; i++ {
		buttonPressTime, speed := i, i
		timeLeft := race.Time - buttonPressTime
		distance := speed * timeLeft
		if distance > race.RecordDistance {
			wins++
		}
	}
	return wins
}

func waysToWin2(race Race) int {
	var wins int
	half := race.Time / 2
	distance := (race.Time - half) * half
	for i := half; distance > race.RecordDistance; i-- {
		distance = (race.Time - i) * i
		if distance > race.RecordDistance {
			wins += 2
		} else {
			break
		}
	}
	if race.Time%2 == 0 {
		wins--
	}

	return wins
}
