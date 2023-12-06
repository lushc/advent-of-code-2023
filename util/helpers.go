package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("parsing number: %s", err)
	}
	return i
}

func ReadInput() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("opening file: %s", err)
	}

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	return lines
}
