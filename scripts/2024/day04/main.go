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

var XMAS = []byte("XMAS")

func part1(input string) (ans int) {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		for j, char := range line {
			if char != 'X' {
				continue
			}

			n := i-3 >= 0
			e := j+3 < len(line)
			s := i+3 < len(lines)
			w := j-3 >= 0

			if n {
				for k, char := range XMAS {
					if char != lines[i-k][j] {
						break
					}
					if char == 'S' {
						ans += 1
					}
				}
			}
			if n && e {
				for k, char := range XMAS {
					if char != lines[i-k][j+k] {
						break
					}
					if char == 'S' {
						ans += 1
					}
				}
			}
			if e {
				for k, char := range XMAS {
					if char != lines[i][j+k] {
						break
					}
					if char == 'S' {
						ans += 1
					}
				}
			}
			if s && e {
				for k, char := range XMAS {
					if char != lines[i+k][j+k] {
						break
					}
					if char == 'S' {
						ans += 1
					}
				}
			}
			if s {
				for k, char := range XMAS {
					if char != lines[i+k][j] {
						break
					}
					if char == 'S' {
						ans += 1
					}
				}
			}
			if s && w {
				for k, char := range XMAS {
					if char != lines[i+k][j-k] {
						break
					}
					if char == 'S' {
						ans += 1
					}
				}
			}
			if w {
				for k, char := range XMAS {
					if char != lines[i][j-k] {
						break
					}
					if char == 'S' {
						ans += 1
					}
				}
			}
			if n && w {
				for k, char := range XMAS {
					if char != lines[i-k][j-k] {
						break
					}
					if char == 'S' {
						ans += 1
					}
				}
			}
		}

	}
	return ans
}

func part2(input string) (ans int) {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if i == 0 || i == len(lines)-1 {
			continue
		}

		for j, char := range line {
			if j == 0 || j == len(line)-1 || char != 'A' {
				continue
			}

			if ((lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S') || (lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M')) &&
				((lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S') || (lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M')) {
				ans += 1
			}
		}
	}

	return ans
}
