package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

	var sum int64
	for _, num := range parsed {
		sum += num
	}

	return sum
}

func part2(input string) int64 {
	parsed := parseInput2(input)

	var sum int64
	for _, num := range parsed {
		sum += num
	}

	return sum
}

func parseInput(input string) (ans []int64) {
	for _, line := range strings.Split(input, "\n") {
		equation := strings.Split(line, ":")
		target, _ := strconv.Atoi(equation[0])
		var nums []int
		fields := strings.Fields(equation[1])
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			nums = append(nums, num)
		}

		maxCombinations := int64(math.Pow(2, float64(len(fields))))
		var i int64

	outer:
		for i = 0; i < maxCombinations; i++ {
			var total int64 = 0
			target64 := int64(target)
			for j, num := range nums {
				if (i & (1 << j)) != 0 {
					total *= int64(num)
				} else {
					total += int64(num)
				}

				if total >= target64 {
					if total == target64 && j == len(nums)-1 {
						ans = append(ans, target64)
						break outer
					}
					continue
				}
			}
		}
	}
	return ans
}

func parseInput2(input string) (ans []int64) {
	for _, line := range strings.Split(input, "\n") {
		equation := strings.Split(line, ":")
		target, _ := strconv.Atoi(equation[0])
		var nums []int
		fields := strings.Fields(equation[1])
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			nums = append(nums, num)
		}

		maxCombinations := int64(math.Pow(3, float64(len(fields))))
		var i int64

	outer:
		for i = 0; i < maxCombinations; i++ {
			var total int64 = 0
			target64 := int64(target)
			for j, num := range nums {
				operator := (i / (int64(math.Pow(3, float64(j))))) % 3

				if operator == 2 {
					total *= int64(num)
				} else if operator == 1 {
					total += int64(num)
				} else if operator == 0 {
					total, _ = strconv.ParseInt(fmt.Sprintf("%v%v", total, num), 10, 64)
					// fmt.Println(total)
				} else {
					panic("unreachable")
				}

				if total >= target64 {
					if total == target64 && j == len(nums)-1 {
						ans = append(ans, target64)
						break outer
					}
					continue
				}
			}
		}
	}
	return ans
}
