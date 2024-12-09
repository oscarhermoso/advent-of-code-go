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
	nums := parseInput(input)

	pos := 0
	head := 0
	headId := 0
	tail := len(nums) - 1
	tailId := len(nums)/2 + 1
	tailSize := 0
	checksum := 0
	debug := ""
	// alreadyDone := true

	for head <= tail {
		fileSize := nums[head]
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

		spaceSize := nums[head]
		head++

		for i := 0; i < spaceSize; i++ {
			if tailSize == 0 {
				if tail <= head {
					break
				}
				tailSize = nums[tail]
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

	return checksum
}

func part2(input string) int {
	nums := parseInput(input)

	pos := 0
	head := 0
	headId := 0
	tail := len(nums) - 1
	tailId := len(nums)/2 + 1
	tailSize := 0
	checksum := 0
	debug := make([]byte, len(nums)*10)

	for i := range debug {
		debug[i] = byte('.')
	}

	for head < len(nums)-1 {
		fileSize := nums[head]
		head++
		for i := 0; i < fileSize; i++ {
			checksum += pos * headId
			debug[pos] = strconv.Itoa(headId)[0]
			pos++
		}
		headId++

		fmt.Println(string(debug))

		spaceSize := nums[head]
		head++

		for i := 0; i < spaceSize; i++ {
			if tailSize == 0 {
				for j := tail; j >= head; j -= 2 {
					if nums[j] != 0 && spaceSize-i >= nums[j] {
						tailSize = nums[j]
						nums[j] = 0
						tailId = j / 2
						break
					} else {
						tailId = 0
					}
				}
			}

			if tailId != 0 {
				debug[pos] = strconv.Itoa(tailId)[0]
				checksum += pos * tailId
				tailSize--
			}
			pos++
		}
	}

	fmt.Println(string(debug))

	return checksum
}

func parseInput(input string) []int {
	var ans []int

	for _, char := range input {
		n, _ := strconv.Atoi(string(char))
		ans = append(ans, n)
	}

	return ans
}
