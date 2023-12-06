package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/lushc/advent-of-code-2023/util"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("opening file: %s", err)
	}

	scanner := bufio.NewScanner(f)
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	calibration := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return
		}

		numbers := make(map[int]string)

		for i, s := range words {
			if first := strings.Index(line, s); first != -1 {
				numbers[first] = fmt.Sprint(i + 1)
			}
			if last := strings.LastIndex(line, s); last != -1 {
				numbers[last] = fmt.Sprint(i + 1)
			}
		}

		for i, char := range line {
			if char < 48 || char > 57 {
				continue
			}
			numbers[i] = string(char)
		}

		keys := make([]int, len(numbers))
		i := 0
		for k := range numbers {
			keys[i] = k
			i++
		}
		sort.Ints(keys)

		calibration += util.ParseInt(numbers[keys[0]] + numbers[keys[len(keys)-1]])
	}

	fmt.Printf("calibration sum is: %d", calibration)
}
