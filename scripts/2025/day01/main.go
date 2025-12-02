package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math"
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
		var dir string
		var amount int
		_, err := fmt.Sscanf(line, "%1s%d", &dir, &amount)
		if err != nil {
			log.Fatalf("failed to parse input: %v", err)
		}
		if dir == "L" {
			amount = -amount
		}
		ans = append(ans, amount)
	}
	return ans
}

func part1(input string) int {
	parsed := parseInput(input)
	total := 50
	passedZero := 0

	for _, rot := range parsed {
		total += rot
		if total%100 == 0 {
			passedZero += 1
		}
	}

	return passedZero
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2(input string) int {
	parsed := parseInput(input)
	total := 50
	passedZero := 0

	for _, rot := range parsed {
		// prevPassedZero := passedZero
		if total%100 == 0 && rot < 0 {
			passedZero -= 1
		}
		prev := int(math.Floor(float64(total) / float64(100)))
		total += rot
		new := int(math.Floor(float64(total) / float64(100)))
		passedZero += Abs(new - prev)
		if total%100 == 0 && rot < 0 {
			passedZero += 1
		}
		// if prevPassedZero < passedZero {
		// 	fmt.Printf("Was rotated %d to point at %d, points at zero %d time(s)\n", rot, total%100)
		// }
	}

	return passedZero
}
