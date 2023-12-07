package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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
}

func part1(lines []string) int {
	players := parseInput(lines)
	sort.Sort(rankSortedPlayers(players))

	var totalWinnings int
	fmt.Println("\nSorted:")
	for i, p := range players {
		totalWinnings += (p.Bid * (i + 1))
		fmt.Println(p, p.handType(), p.Bid, i+1)
	}

	return totalWinnings
}

type Player struct {
	Hand string
	Bid  int
}

type rankSortedPlayers []Player

func (p rankSortedPlayers) Len() int {
	return len(p)
}

func (p rankSortedPlayers) Less(i, j int) bool {
	lesserHandType := Player.handType(p[i]) < Player.handType(p[j])
	equalHandType := Player.handType(p[i]) == Player.handType(p[j])
	return lesserHandType || (equalHandType && lowerSingleCard(p[i], p[j]))
}

func (p rankSortedPlayers) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (p Player) handType() int {
	cardMap := make(map[rune]int)
	for _, c := range p.Hand {
		cardMap[c]++
	}

	if anyInMap(cardMap, 5) {
		return FiveOfAKind
	} else if anyInMap(cardMap, 4) {
		return FourOfAKind
	} else if anyInMap(cardMap, 3) {
		if anyInMap(cardMap, 2) {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	} else if anyInMap(cardMap, 2) {
		if valueCountInMap(cardMap, 1) == 1 {
			return TwoPair
		} else {
			return OnePair
		}
	}
	return HighCard
}

func lowerSingleCard(p1, p2 Player) bool {
	cards := []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	cardValueMap := make(map[byte]int)
	for i := 0; i < len(cards); i++ {
		cardValueMap[cards[i]] = i
	}
	fmt.Println(cardValueMap)

	for i := 0; i < len(p1.Hand); i++ {
		v1, v2 := cardValueMap[p1.Hand[i]], cardValueMap[p2.Hand[i]]
		if v1 == v2 {
			continue
		}
		if v1 < v2 {
			return true
		} else {
			return false
		}
	}
	return false
}

func anyInMap(m map[rune]int, val int) bool {
	for _, v := range m {
		if v == val {
			return true
		}
	}

	return false
}

func valueCountInMap(m map[rune]int, val int) int {
	var count int
	for _, v := range m {
		if v == val {
			count++
		}
	}

	return count
}

func parseInput(in []string) []Player {
	var players []Player
	for _, line := range in {
		p := strings.Fields(line)
		bid, err := strconv.Atoi(p[1])
		if err != nil {
			log.Fatal(err)
		}
		players = append(players, Player{
			Hand: p[0],
			Bid:  bid,
		})
	}

	return players
}

func betterHand(p1, p2 Player) bool {

	return false
}
