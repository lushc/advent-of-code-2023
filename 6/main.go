package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/lushc/advent-of-code-2023/util"
)

var reNumbers = regexp.MustCompile(`\d+`)

func main() {
	times := []string{}
	distances := []string{}

	for i, line := range util.ReadInput() {
		if line == "" {
			continue
		}
		tokens := reNumbers.FindAllString(line, -1)
		if i == 0 {
			times = append(times, tokens...)
		} else {
			distances = append(distances, tokens...)
		}
	}

	partOne(times, distances)
	partTwo(times, distances)
}

func partOne(times, distances []string) {
	start := time.Now()
	ways := 1
	for i, time := range times {
		ways *= searchWays(util.ParseInt(time), util.ParseInt(distances[i]))
	}
	elapsed := time.Since(start)
	fmt.Printf("part 1: total ways to win is %d (took %s)\n", ways, elapsed)
}

func partTwo(times, distances []string) {
	start := time.Now()
	ways := searchWays(
		util.ParseInt(strings.Join(times, "")),
		util.ParseInt(strings.Join(distances, "")),
	)
	elapsed := time.Since(start)
	fmt.Printf("part 2: total ways to win is %d (took %s)\n", ways, &elapsed)
}

func searchWays(time, distance int) int {
	min := 0
	for speed := 0; speed < time; speed++ {
		dist := speed * (time - speed)
		if dist > distance {
			min = speed
			break
		}
	}

	max := 0
	for speed := time; speed > min; speed-- {
		dist := speed * (time - speed)
		if dist > distance {
			max = speed
			break
		}
	}

	return max - min + 1
}
