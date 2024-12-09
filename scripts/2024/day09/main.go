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
	parsed := parseInput(input)
	return parsed
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

func parseInput(input string) int {
	var ans []int

	for _, char := range input {
		n, _ := strconv.Atoi(string(char))
		ans = append(ans, n)
	}

	pos := 0
	head := 0
	headId := 0
	tail := len(ans) - 1
	tailId := len(ans)/2 + 1
	tailSize := 0
	checksum := 0
	debug := ""
	// alreadyDone := true

	for head <= tail {
		fileSize := ans[head]
		head++

		for i := 0; i < fileSize; i++ {
			checksum += pos * headId
			pos++
			debug = fmt.Sprintf("%s%d", debug, headId)
		}
		headId++

		if head == tail { // does nothign!
			break
		}

		spaceSize := ans[head]
		head++

		for i := 0; i < spaceSize; i++ {
			if tail < head-2 {
				break
			} else if tailSize == 0 {
				tailSize = ans[tail]
				tail -= 2
				tailId--
			}
			checksum += pos * tailId
			pos++
			tailSize--
			debug = fmt.Sprintf("%s%d", debug, tailId)
		}
	}

	for i := 0; i < tailSize; i++ {
		checksum += pos * tailId
		debug = fmt.Sprintf("%s%d", debug, tailId)
	}

	if debug == "0099811188827773336446555566" {
		print("input matched!\n")
	} else if len(debug) < 1000 {
		fmt.Printf("RUH ROH, no matchy")
		fmt.Printf("should be:\n0099811188827773336446555566\n")
		fmt.Printf("but was:\n%s\n", debug)
	}

	return checksum
}
