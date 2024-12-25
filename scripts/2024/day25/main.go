package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	locks, keys := parseInput(input)
	ans := 0

	for _, lock := range locks {
	keyLoop:
		for _, key := range keys {
			for i := 0; i < 5; i++ {
				if key[i]+lock[i] > 5 {
					continue keyLoop
				}
			}
			ans += 1
			// fmt.Printf("lock: %v, key: %v\n", lock, key)
		}
	}

	return ans
}

func part2(input string) int {
	keys, locks := parseInput(input)
	_, _ = locks, keys

	return 0
}

func parseInput(input string) (locks [][]int, keys [][]int) {

	schematics := strings.Split(input, "\n\n")

	for _, schematic := range schematics {
		height := []int{0, 0, 0, 0, 0}
		for _, line := range strings.Split(schematic, "\n")[1:6] {
			for j, char := range line {
				if char == '#' {
					height[j] += 1
				}
			}
		}

		if schematic[0] == '#' {
			locks = append(locks, height)
		} else {
			keys = append(keys, height)
		}
	}

	return locks, keys
}
