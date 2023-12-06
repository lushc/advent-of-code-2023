package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/lushc/advent-of-code-2023/util"
)

type Coordinate struct {
	x int
	y int
}

type GridNumber struct {
	value       int
	topLeft     Coordinate
	bottomRight Coordinate
}

type GridSymbol struct {
	value      string
	adjacentTo []GridNumber
}

var (
	reNumbers = regexp.MustCompile(`\d+`)
	reSymbols = regexp.MustCompile(`[^\d\.]`)
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("opening file: %s", err)
	}

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return
		}
		lines = append(lines, line)
	}

	numbers := findGridNumbers(lines)
	symbols := findSymbols(lines)

	partSum := 0
	for _, number := range numbers {
		for y := number.topLeft.y; y <= number.bottomRight.y; y++ {
			for x := number.topLeft.x; x <= number.bottomRight.x; x++ {
				symbol, ok := symbols[Coordinate{x: x, y: y}]
				if !ok {
					continue
				}
				partSum += number.value

				if symbol.value == "*" {
					symbol.adjacentTo = append(symbol.adjacentTo, number)
				}
			}
		}
	}

	ratioSum := 0
	for _, symbol := range symbols {
		if symbol.value == "*" && len(symbol.adjacentTo) == 2 {
			ratioSum += symbol.adjacentTo[0].value * symbol.adjacentTo[1].value
		}
	}

	fmt.Printf("part sum: %d\nratio sum: %d", partSum, ratioSum)
}

func findGridNumbers(lines []string) []GridNumber {
	numbers := []GridNumber{}
	for y, line := range lines {
		nums := reNumbers.FindAllStringSubmatch(line, -1)
		numIndices := reNumbers.FindAllStringSubmatchIndex(line, -1)

		for i, v := range nums {
			top := Coordinate{x: numIndices[i][0], y: y}
			bottom := Coordinate{x: numIndices[i][1], y: y}

			if y > 0 {
				top.y -= 1
			}
			if y+1 < len(lines) {
				bottom.y += 1
			}
			if top.x > 0 {
				top.x -= 1
			}

			numbers = append(numbers, GridNumber{
				value:       util.ParseInt((v[0])),
				topLeft:     top,
				bottomRight: bottom,
			})
		}
	}
	return numbers
}

func findSymbols(lines []string) map[Coordinate]*GridSymbol {
	symbols := map[Coordinate]*GridSymbol{}
	for y, line := range lines {
		match := reSymbols.FindAllStringSubmatch(line, -1)
		matchIndices := reSymbols.FindAllStringSubmatchIndex(line, -1)

		for i, v := range match {
			coord := Coordinate{x: matchIndices[i][0], y: y}
			symbols[coord] = &GridSymbol{
				value: v[0],
			}
		}
	}
	return symbols
}
