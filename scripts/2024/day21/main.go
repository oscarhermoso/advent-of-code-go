package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image"
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

// var numpad [][]rune = [][]rune{
// 	{'7', '8', '9'},
// 	{'4', '5', '6'},
// 	{'1', '2', '3'},
// 	{' ', '0', 'A'},
// }

// var dirpad [][]rune = [][]rune{
// 	{' ', '^', 'A'},
// 	{'<', 'V', '>'},
// }

var numpad map[rune]image.Point = map[rune]image.Point{
	'7': {0, 3}, '8': {1, 3}, '9': {2, 3},
	'4': {0, 2}, '5': {1, 2}, '6': {2, 2},
	'1': {0, 1}, '2': {1, 1}, '3': {2, 1},
	/*		  */ '0': {1, 0}, 'A': {2, 0},
}

var dirpad map[rune]image.Point = map[rune]image.Point{
	/*		  */ '^': {1, 1}, 'A': {2, 1},
	'<': {0, 0}, 'v': {1, 0}, '>': {2, 0},
}

func part1(input string) int {
	codes := parseInput(input)
	_ = codes

	ans := 0

	for _, code := range codes {
		pressureLoc, radLoc, coldLoc := 'A', 'A', 'A'
		path := code
		path, pressureLoc = getPath(pressureLoc, path, numpad)
		path, radLoc = getPath(radLoc, path, dirpad)
		path, coldLoc = getPath(coldLoc, path, dirpad)
		// code, hisLoc = getPath(hisLoc, code, dirpad)

		num, _ := strconv.Atoi(string(code[0:3]))
		ans += num * len(path)

		fmt.Printf("%s: %s\n", string(code), string(path))
	}

	return ans
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

func parseInput(input string) (codes [][]rune) {
	for _, line := range strings.Split(input, "\n") {
		var code []rune
		for _, char := range line {
			code = append(code, char)
		}
		codes = append(codes, code)
	}
	return codes
}

func getPath(start rune, input []rune, keypad map[rune]image.Point) (output []rune, end rune) {
	pos := keypad[start]

	for _, key := range input {
		dest := keypad[key]
		delta := dest.Sub(pos)
		// fmt.Printf("uhhhh: from %v to %v\n", pos, dest)
		if start == 'A' && key == 'v' {
			output = append(output, '<')
			output = append(output, 'v')
		} else if start == '<' || start == '>' {
			if delta.X < 0 {
				output = append(output, []rune(strings.Repeat("<", -delta.X))...)
			} else if delta.X > 0 {
				output = append(output, []rune(strings.Repeat(">", delta.X))...)
			}
			if delta.Y < 0 {
				output = append(output, []rune(strings.Repeat("v", -delta.Y))...)
			} else if delta.Y > 0 {
				output = append(output, []rune(strings.Repeat("^", delta.Y))...)
			}
		} else {
			if delta.Y < 0 {
				output = append(output, []rune(strings.Repeat("v", -delta.Y))...)
			} else if delta.Y > 0 {
				output = append(output, []rune(strings.Repeat("^", delta.Y))...)
			}
			if delta.X < 0 {
				output = append(output, []rune(strings.Repeat("<", -delta.X))...)
			} else if delta.X > 0 {
				output = append(output, []rune(strings.Repeat(">", delta.X))...)
			}
		}
		output = append(output, 'A')
		start = key
		pos = dest
	}
	// fmt.Printf("%s\n", string(output))
	return output, end
}
