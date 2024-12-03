package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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

func part1(input string) (ans int) {
	parsed := parseInput(input)
	
	for _, mult := range parsed {
		ans += mult
	}

	return ans
}

func part2(input string) (ans int) {
	var doLines string

	for _, do := range strings.Split(input, `do()`) {
		dontIndex := strings.Index(do, "don't()")
		if dontIndex == -1 {
			doLines += do
		} else {
			doLines += do[:dontIndex]
		}
	}

	parsed := parseInput(string(doLines))

	for _, mult := range parsed {
		ans += mult
	}
	
	return ans
}

func parseInput(input string) (ans []int) {
	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	for _, line := range strings.Split(input, "\n") {
		lineAns := 0
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			l, _ := strconv.Atoi(match[1])
			r, _ := strconv.Atoi(match[2])
			lineAns += (l*r)
		}
		ans = append(ans, lineAns)
	}
	return ans
}
