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

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		var bank []int
		for _, char := range line {
			joltage, _ := strconv.Atoi(string(char))
			bank = append(bank, joltage)
		}
		ans = append(ans, bank)
	}
	return ans
}

func part1(input string) int {
	banks := parseInput(input)
	var combinedMax int

	for _, bank := range banks {
		var max int
		for i, bat1 := range bank[:len(bank)-1] {
			for _, bat2 := range bank[i+1:] {
				total := bat1*10 + bat2
				if total > max {
					// fmt.Printf("found new max, %d (%d) + %d (%d) = %d \n", i, bat1, j, bat2, max)
					max = total
				}
			}
		}
		fmt.Printf("best combo %d\n", max)
		combinedMax += max
		max = 0
	}

	return combinedMax
}

func part2(input string) int64 {
	banks := parseInput(input)

	var totalJoltage int64

	for _, bank := range banks {
		joltageStr := recurse(bank, 12, "")
		joltage, _ := strconv.ParseInt(joltageStr, 10, 64)
		totalJoltage += joltage
	}

	return totalJoltage
}

func recurse(input []int, remaining int, result string) string {
	if remaining == 0 {
		return result
	}

	var max int

	for _, char := range input[:len(input)-remaining+1] {
		if char > max {
			max = char
		}
	}

	for i, char := range input[:len(input)-remaining+1] {
		if char == max {
			result += fmt.Sprint(char)
			// fmt.Printf("found best char %d at pos %d. result is %s with %v remaining\n", max, i, result, input[i+1:])
			return recurse(input[i+1:], remaining-1, result)
		}
	}

	log.Fatalf("recurse: unreachable")
	return ""
}
