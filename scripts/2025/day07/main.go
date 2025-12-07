package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
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

func parseInput(input string) (ans []int) {
	for _, line := range strings.Split(input, "\n") {
		var arg1, arg2 int
		_, err := fmt.Sscanf(line, "<format>", &arg1, &arg2)
		if err != nil {
			log.Fatalf("failed to parse input: %v", err)
		}
		ans = append(ans, arg1, arg2)
	}
	return ans
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
