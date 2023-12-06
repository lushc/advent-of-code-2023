package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/lushc/advent-of-code-2023/util"
)

const (
	Red   = "red"
	Green = "green"
	Blue  = "blue"
)

type Set map[string]int

func main() {
	re := regexp.MustCompile(`(\d+)\s(red|green|blue)`)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("opening file: %s", err)
	}

	bag := Set{Red: 12, Green: 13, Blue: 14}
	idSum := 0
	powerSum := 0

	i := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i++
		line := scanner.Text()
		if line == "" {
			return
		}

		matches := re.FindAllStringSubmatch(line, -1)
		minimum := Set{Red: 0, Green: 0, Blue: 0}
		valid := true
		for _, v := range matches {
			colour := v[2]
			count := util.ParseInt(v[1])
			if count > bag[colour] {
				valid = false
			}
			if count > minimum[colour] {
				minimum[colour] = count
			}
		}

		if valid {
			idSum += i
		}

		powerSum += (minimum[Red] * minimum[Green] * minimum[Blue])
	}

	fmt.Printf("sum of IDs: %d\n", idSum)
	fmt.Printf("sum of powers: %d\n", powerSum)
}
