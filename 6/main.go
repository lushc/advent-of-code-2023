package main

import (
	"fmt"
	"regexp"
	"strings"

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

	partOneWays := 1
	for i, time := range times {
		partOneWays *= calculateWays(util.ParseInt(time), util.ParseInt(distances[i]))
	}
	fmt.Printf("part 1: total ways to win is %d\n", partOneWays)

	time := util.ParseInt(strings.Join(times, ""))
	distance := util.ParseInt(strings.Join(distances, ""))
	partTwoWays := calculateWays(time, distance)
	fmt.Printf("part 2: total ways to win is %d\n", partTwoWays)

}

func calculateWays(time, distance int) int {
	min := 0
	for speed := 0; speed < time; speed++ {
		dist := speed * (time - speed)
		if dist > distance {
			min = speed
			break
		}
	}
	fmt.Printf("min speed for time %d is %d\n", time, min)

	max := 0
	for speed := time; speed > min; speed-- {
		dist := speed * (time - speed)
		if dist > distance {
			max = speed
			break
		}
	}
	fmt.Printf("max speed for time %d is %d\n", time, max)

	return max - min + 1
}
