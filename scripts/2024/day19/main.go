package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
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
	towels, designs := parseInput(input)
	ans := 0
	cache := make(map[string]bool)
	var dfs func(design string) bool
	dfs = func(design string) bool {
		if len(design) == 0 {
			return true
		}
		if success, ok := cache[design]; ok {
			return success
		}
		for _, towel := range towels {
			if strings.HasPrefix(design, towel) {
				success := dfs(design[len(towel):])
				cache[design] = success
				if success {
					return true
				}
			}
		}
		return false
	}
	for _, design := range designs {
		// fmt.Printf("Checking design %s (%d/%d)\n", design, i, len(designs)-1)
		found := dfs(design)
		if found {
			ans += 1
		}
	}
	return ans
}

func part2(input string) int {
	towels, designs := parseInput(input)
	ans := 0

	cache := make(map[string]int)

	var dfs func(design string) int
	dfs = func(design string) int {
		if len(design) == 0 {
			return 1
		}
		if count, ok := cache[design]; ok {
			return count
		}
		count := 0
		for _, towel := range towels {
			if strings.HasPrefix(design, towel) {
				count += dfs(design[len(towel):])
			}
		}
		cache[design] = count
		return count
	}

	for _, design := range designs {
		// fmt.Printf("Checking design %s (%d/%d)\n", design, i, len(designs)-1)
		ans += dfs(design)
	}

	return ans
}

func parseInput(input string) (towels []string, designs []string) {
	lines := strings.Split(input, "\n")

	towels = strings.Split(lines[0], ", ")
	slices.SortFunc(towels, func(a string, b string) int {
		if len(a) == len(b) {
			return 0
		}
		return -(len(a) - len(b)) // desc by length
	})

	designs = lines[2:]

	return towels, designs
}
