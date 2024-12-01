package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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
	left, right := parseInput(input)

	ans := 0

	for i := range left {
		ans += int(math.Abs(float64(left[i] - right[i])))
	}

	return ans
}

func part2(input string) int {
	left, right := parseInput(input)

	ans := 0

	for _, l := range left {
		multiplier := 0
		for _, r := range right {
			if l == r {
				multiplier += 1
			}
		}
		ans += l * multiplier
	}

	return ans
}

func parseInput(input string) (left []int, right []int) {
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)

		l, _ := strconv.Atoi(fields[0])
		r, _ := strconv.Atoi(fields[1])

		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	return left, right
}
