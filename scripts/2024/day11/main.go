package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
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
	_ = parsed

	return 0
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

func parseInput(input string) (ans []int) {
	for _, line := range strings.Split(input, "\n") {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("error converting string to int: %v", err)
		}
		ans = append(ans, val)
	}
	return ans
}
