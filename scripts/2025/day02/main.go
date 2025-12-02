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
	// input = strings.TrimRight(input, "\n")
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

func parseInput(input string) (ans [][]string) {
	for _, productRange := range strings.Split(input, ",") {
		pair := strings.Split(productRange, "-")
		if len(pair[0])%2 == 1 {
			pair[0] = "0" + pair[0]
		}
		if len(pair[1])%2 == 1 {
			pair[1] = "0" + pair[1]
		}
		if len(pair[1])-len(pair[0]) == 0 {
			ans = append(ans, pair)
		} else if len(pair[1])-len(pair[0]) == 2 {
			ans = append(ans, []string{"00" + pair[0], pair[1]})
		} else {
			log.Fatalf("Unexpected string length diff %d for strings: %s, %s", len(pair[1])-len(pair[0]), pair[0], pair[1])
		}
	}
	return ans
}

func part1(input string) int64 {
	parsed := parseInput(input)
	var total int64
	for _, pair := range parsed {
		halfStr := pair[0][:len(pair[0])/2]
		halfInt, _ := strconv.ParseInt(halfStr, 10, 64)
		min, _ := strconv.ParseInt(pair[0], 10, 64)
		max, _ := strconv.ParseInt(pair[1], 10, 64)

		wholeStr := halfStr + halfStr
		wholeInt, _ := strconv.ParseInt(wholeStr, 10, 64)
		for wholeInt <= max {
			if wholeInt >= min {
				fmt.Printf("found invalid ID: %s in range %s-%s\n", wholeStr, pair[0], pair[1])
				total += wholeInt
			}
			halfInt += 1
			halfStr = strconv.FormatInt(halfInt, 10)
			wholeStr = halfStr + halfStr
			wholeInt, _ = strconv.ParseInt(wholeStr, 10, 64)
		}
	}
	return total
}

func part2(input string) int64 {
	parsed := parseInput(input)
	_ = parsed

	return 0
}
