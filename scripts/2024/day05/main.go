package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)

	sum := 0

	for _, val := range parsed {
		sum += val
	}

	return sum
}

func part2(input string) int {
	parsed := parseInput2(input)

	sum := 0

	for _, val := range parsed {
		sum += val
	}

	return sum
}

func parseInput(input string) (ans []int) {
	afters := make(map[string][]string)
	lines := strings.Split(input, "\n")
	part2 := 0

	for i, line := range lines {
		if line == "" {
			part2 = i
			break
		}

		pages := strings.Split(line, "|")

		afters[pages[0]] = append(afters[pages[0]], pages[1])
	}

lineLoop:
	for _, line := range lines[part2+1:] {
		pages := strings.Split(line, ",")

		for j, page := range pages {
			befores := pages[:j]
			for _, before := range befores {
				if slices.Contains(afters[page], before) {
					// fmt.Printf("page %v should not be before %v but was (%v)\n", before, page, pages)
					continue lineLoop
				}
			}
		}

		middleIndex := len(pages) / 2
		middleVal, _ := strconv.Atoi(pages[middleIndex])

		ans = append(ans, middleVal)
	}

	return ans
}

func parseInput2(input string) (ans []int) {
	afters := make(map[string][]string)
	lines := strings.Split(input, "\n")
	part2 := 0

	for i, line := range lines {
		if line == "" {
			part2 = i
			break
		}

		pages := strings.Split(line, "|")

		afters[pages[0]] = append(afters[pages[0]], pages[1])
	}

	var unsortedLines [][]string

lineLoop:
	for _, line := range lines[part2+1:] {
		pages := strings.Split(line, ",")

		for j, page := range pages {
			befores := pages[:j]
			for _, before := range befores {
				if slices.Contains(afters[page], before) {
					unsortedLines = append(unsortedLines, pages)
					continue lineLoop
				}
			}
		}
	}

	for _, pages := range unsortedLines {
		for {
			swapped := false

		pageLoop:
			for j, page := range pages {
				befores := pages[:j]
				for _, before := range befores {
					if slices.Contains(afters[page], before) {
						pages[j] = pages[j-1]
						pages[j-1] = page

						swapped = true
						continue pageLoop
					}
				}
			}

			if !swapped {
				middleIndex := len(pages) / 2
				middleVal, _ := strconv.Atoi(pages[middleIndex])
				ans = append(ans, middleVal)
				break
			}
		}
	}

	return ans
}
