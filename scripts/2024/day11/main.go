package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	var stones, nextStones []string
	stones = strings.Fields(input)

	for i := 0; i < 25; i++ {
		for _, stone := range stones {
			if stone == "0" {
				nextStones = append(nextStones, "1")
			} else if len(stone)%2 == 0 {
				stone0, _ := strconv.ParseInt(stone[:len(stone)/2], 10, 64)
				stone1, _ := strconv.ParseInt(stone[len(stone)/2:], 10, 64)
				nextStones = append(nextStones, fmt.Sprint(stone0))
				nextStones = append(nextStones, fmt.Sprint(stone1))
			} else {
				stone0, _ := strconv.ParseInt(stone, 10, 64)
				nextStones = append(nextStones, fmt.Sprint(stone0*2024))
			}
		}
		// fmt.Println(nextStones)
		stones = nextStones
		nextStones = []string{}
	}

	return len(stones)
}

func part2(input string) int {
	// parsed := parseInput(input)
	// _ = parsed

	return 0
}
