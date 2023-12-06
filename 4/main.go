package main

import (
	"fmt"
	"strings"

	"github.com/lushc/advent-of-code-2023/util"
)

type Card struct {
	have    []int
	winning []int
}

func main() {
	cards := [][]Card{}
	lines := util.ReadInput()
	for _, line := range lines {
		card := Card{}
		var slice *[]int

		for _, token := range strings.Split(line, " ") {
			if token == "" {
				continue
			}
			if strings.HasSuffix(token, ":") {
				slice = &card.have
				continue
			}
			if token == "|" {
				slice = &card.winning
				continue
			}
			if slice != nil {
				*slice = append(*slice, util.ParseInt(token))
			}
		}

		cards = append(cards, []Card{card})
	}

	worthSum := 0
	for _, card := range cards {
		if points := card[0].Points(); points > 0 {
			worthSum += points
		}
	}

	totalCards := len(cards)
	for i, set := range cards {
		for _, card := range set {
			matches := card.Matches()
			if matches == 0 {
				break
			}
			for j := 1; j <= matches; j++ {
				next := i + j
				if next < len(cards) {
					cards[next] = append(cards[next], cards[next][0])
				}
			}
			totalCards += matches
		}
	}

	fmt.Printf("total points worth: %d\n", worthSum)
	fmt.Printf("total scratchcards: %d\n", totalCards)
}

func (c Card) Points() int {
	points := 0
	for _, got := range c.have {
		for _, winner := range c.winning {
			if got != winner {
				continue
			}
			if points == 0 {
				points += 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func (c Card) Matches() int {
	matches := 0
	for _, got := range c.have {
		for _, winner := range c.winning {
			if got == winner {
				matches += 1
			}
		}
	}
	return matches
}
