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
		var nums []int64
		fields := strings.Fields(equation[1])
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			nums = append(nums, int64(num))
		}
		if dfs(int64(target), false, nums...) {
			ans = append(ans, int64(target))
		}
	}
	return ans
}

func parseInput2(input string) (ans []int64) {
	for _, line := range strings.Split(input, "\n") {
		equation := strings.Split(line, ":")
		target, _ := strconv.Atoi(equation[0])
		var nums []int64
		fields := strings.Fields(equation[1])
		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			nums = append(nums, int64(num))
		}
		if dfs(int64(target), true, nums...) {
			ans = append(ans, int64(target))
		}
	}
	return ans
}

func dfs(target int64, part2 bool, nums ...int64) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}
	if nums[0] > target {
		return false
	}

	if success := dfs(target, part2, append([]int64{nums[0] * nums[1]}, nums[2:]...)...); success {
		return success
	}
	if success := dfs(target, part2, append([]int64{nums[0] + nums[1]}, nums[2:]...)...); success {
		return success
	}
	if part2 {
		numNum, _ := strconv.ParseInt(fmt.Sprintf("%v%v", nums[0], nums[1]), 10, 64)
		return dfs(target, part2, append([]int64{numNum}, nums[2:]...)...)
	}

	return false
}
