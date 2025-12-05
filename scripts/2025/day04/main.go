package main

import (
	_ "embed"
	"flag"
	"fmt"
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

	if part >= 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}

func part1(input string) (count int) {
	parsed := parseInput(input)
	var rollCount int
	var inaccessible int

	for y, line := range parsed {
		for x, char := range line {

			if char != '@' {
				continue
			}

			rollCount += 1

			var adj int

			if x != 0 {
				if line[x-1] != '.' {
					// fmt.Println("found W")
					adj += 1
				}

				if y != 0 && parsed[y-1][x-1] != '.' {
					// fmt.Println("found NW")
					adj += 1
				}
				if y != len(parsed)-1 && parsed[y+1][x-1] != '.' {
					// fmt.Println("found SW")
					adj += 1
				}
			}

			if x != len(line)-1 {
				if line[x+1] != '.' {
					// fmt.Println("found E")
					adj += 1
				}

				if y != 0 && parsed[y-1][x+1] != '.' {
					// fmt.Println("found NE")
					adj += 1
				}
				if y != len(parsed)-1 && parsed[y+1][x+1] != '.' {
					// fmt.Println("found SE")
					adj += 1
				}
			}

			if y != 0 && parsed[y-1][x] != '.' {
				// fmt.Println("found N")
				adj += 1
			}
			if y != len(parsed)-1 && parsed[y+1][x] != '.' {
				// fmt.Println("found S")
				adj += 1
			}

			fmt.Printf("tp %d,%d adj had %d adjacent\n", x, y, adj)
			if adj >= 4 {
				// fmt.Println("found inaccessible?")
				inaccessible += 1
				continue
			}
		}
	}

	return rollCount - inaccessible
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}
