package main

import (
	"fmt"
	"strings"

	"github.com/lushc/advent-of-code-2023/util"
)

type Map struct {
	entries []MapEntry
}

type MapEntry struct {
	destStart int
	srcStart  int
	length    int
}

type SeedRange struct {
	start  int
	length int
}

func main() {
	seeds := []int{}
	seedRanges := []SeedRange{}
	// 0: seed-to-soil
	// 1: soil-to-fertilizer
	// 2: fertilizer-to-water
	// 3: water-to-light
	// 4: light-to-temperature
	// 5: temperature-to-humidity
	// 6: humidity-to-location
	maps := make([]Map, 7)

	parseMapIndex := -1
	lines := util.ReadInput()
	for i, line := range lines {
		if line == "" {
			continue
		}

		// first line is the seeds
		if i == 0 {
			tokens := strings.Split(line, " ")[1:]
			// get all seeds for part 1
			for _, token := range tokens {
				seeds = append(seeds, util.ParseInt(token))
			}
			// get seed ranges for part 2
			for next := 0; next < len(tokens); next += 2 {
				pair := tokens[next : next+2]
				seedRanges = append(seedRanges, SeedRange{
					start:  util.ParseInt(pair[0]),
					length: util.ParseInt(pair[1]),
				})
			}
			continue
		}

		// remaining lines are map names and their values
		tokens := strings.Split(line, " ")
		if len(tokens) == 2 {
			parseMapIndex++
			continue
		}

		maps[parseMapIndex].entries = append(maps[parseMapIndex].entries, MapEntry{
			destStart: util.ParseInt(tokens[0]),
			srcStart:  util.ParseInt(tokens[1]),
			length:    util.ParseInt(tokens[2]),
		})
	}

	lowestLocations := []int{-1, -1}
	for _, seed := range seeds {
		next := seed
		for _, m := range maps {
			next = m.Lookup(next)
		}
		if lowestLocations[0] == -1 || next < lowestLocations[0] {
			lowestLocations[0] = next
		}
	}
	// this is really really slow :')
	for _, seedRange := range seedRanges {
		for i := seedRange.start; i < seedRange.start+seedRange.length; i++ {
			next := i
			for _, m := range maps {
				next = m.Lookup(next)
			}
			if lowestLocations[1] == -1 || next < lowestLocations[1] {
				lowestLocations[1] = next
			}
		}
	}

	fmt.Printf("part 1 lowest location is: %d\n", lowestLocations[0])
	fmt.Printf("part 2 lowest location is: %d\n", lowestLocations[1])
}

func (m Map) Lookup(src int) int {
	for _, entry := range m.entries {
		// source must be in range to create a valid offset
		if src >= entry.srcStart && src <= entry.srcStart+(entry.length-1) {
			return entry.destStart + (src - entry.srcStart)
		}
	}
	return src
}
