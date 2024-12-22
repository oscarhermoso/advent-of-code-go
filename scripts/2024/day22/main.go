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

func part1(input string) int64 {
	parsed := parseInput(input)
	sum := int64(0)

	for _, val := range parsed {
		sum += val
	}

	return sum
}

func part2(input string) int64 {
	parsed := parseInput2(input)
	max := int64(0)

	for _, val := range parsed {
		if val > max {
			// fmt.Printf("best seq is %s, with value %d\n", seq, val)
			max = val
		}
	}

	return max
}

func parseInput(input string) (ans []int64) {
	for _, line := range strings.Split(input, "\n") {
		val, _ := strconv.ParseInt(line, 10, 64)

		for i := 0; i < 2000; i++ {
			val = ((val * 64) ^ val) % 16777216
			val = ((val / 32) ^ val) % 16777216
			val = ((val * 2048) ^ val) % 16777216
		}

		ans = append(ans, val)
	}
	return ans
}

func parseInput2(input string) map[string]int64 {
	totals := make(map[string]int64)

	for _, line := range strings.Split(input, "\n") {
		val, _ := strconv.ParseInt(line, 10, 64)

		var prev3, prev2, prev1, prev0 int64
		prev0 = val % 10
		sequences := make(map[string]bool)

		for i := 0; i < 2000; i++ {
			val = ((val * 64) ^ val) % 16777216
			val = ((val / 32) ^ val) % 16777216
			val = ((val * 2048) ^ val) % 16777216

			price := val % 10

			// fmt.Printf("%8d: %d, (%d)\n", val, price, price-prev0)

			if i > 2 {
				seq := fmt.Sprintf("%d%d%d%d", prev2-prev3, prev1-prev2, prev0-prev1, price-prev0)
				_, found := sequences[seq]

				if !found {
					sequences[seq] = true
					totals[seq] += price
					// if seq == "-99-10" {
					// 	initial, _ := strconv.ParseInt(line, 10, 64)
					// 	fmt.Printf("with initial %d, selling for price %d when sequence is %s\n", initial, price, seq)
					// }
				}
			}
			prev3, prev2, prev1, prev0 = prev2, prev1, prev0, price
		}
	}

	return totals
}
