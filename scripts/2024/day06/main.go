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

func part1(input string) int {
	start, _ := parseInput(input)
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])

	x := start[0]
	y := start[1]
	dir := '^'

	visited := 0

	for 0 <= x && x < width && 0 <= y && y < height {
		if lines[y][x] != 'X' {
			line := []rune(lines[y])
			line[x] = 'X'
			lines[y] = string(line)
			visited += 1
		}

		if dir == '^' {
			if y-1 >= 0 && lines[y-1][x] == '#' {
				dir = '>'
			} else {
				y -= 1
			}
		} else if dir == '>' {
			if x+1 < width && lines[y][x+1] == '#' {
				dir = 'v'
			} else {
				x += 1
			}
		} else if dir == 'v' {
			if y+1 < height && lines[y+1][x] == '#' {
				dir = '<'
			} else {
				y += 1
			}
		} else if dir == '<' {
			if x-1 >= 0 && lines[y][x-1] == '#' {
				dir = '^'
			} else {
				x -= 1
			}
		}
	}

	return visited
}

func part2(input string) int {
	start, _ := parseInput(input)
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])

	x := start[0]
	y := start[1]
	dir := '^'

	for 0 <= x && x < width && 0 <= y && y < height {
		if lines[y][x] != 'X' {
			line := []rune(lines[y])
			line[x] = 'X'
			lines[y] = string(line)
		}

		if dir == '^' {
			if y-1 >= 0 && lines[y-1][x] == '#' {
				dir = '>'
			} else {
				y -= 1
			}
		} else if dir == '>' {
			if x+1 < width && lines[y][x+1] == '#' {
				dir = 'v'
			} else {
				x += 1
			}
		} else if dir == 'v' {
			if y+1 < height && lines[y+1][x] == '#' {
				dir = '<'
			} else {
				y += 1
			}
		} else if dir == '<' {
			if x-1 >= 0 && lines[y][x-1] == '#' {
				dir = '^'
			} else {
				x -= 1
			}
		}
	}

	blocks := 0

	for yPos, line := range lines {
		for xPos, loc := range line {
			if (loc == 'X') && !(xPos == start[0] && yPos == start[1]) {
				line := []rune(lines[yPos])
				line[xPos] = '#'
				lines[yPos] = string(line)

				visisted := make(map[string]bool)
				x = start[0]
				y = start[1]
				dir = '^'

				for 0 <= x && x < width && 0 <= y && y < height {
					if dir == '^' {
						if y-1 >= 0 && lines[y-1][x] == '#' {
							dir = '>'
							loc := fmt.Sprintf("%v|%v|%v", x, y, dir)
							v := visisted[loc]
							if v {
								blocks += 1
								break
							}
							visisted[loc] = true

						} else {
							y -= 1
						}
					} else if dir == '>' {
						if x+1 < width && lines[y][x+1] == '#' {
							dir = 'v'
						} else {
							x += 1
						}
					} else if dir == 'v' {
						if y+1 < height && lines[y+1][x] == '#' {
							dir = '<'
						} else {
							y += 1
						}
					} else if dir == '<' {
						if x-1 >= 0 && lines[y][x-1] == '#' {
							dir = '^'
						} else {
							x -= 1
						}
					}
				}

				line = []rune(lines[yPos])
				line[xPos] = loc
				lines[yPos] = string(line)
			}
		}
	}

	return blocks
}

func parseInput(input string) (start []int, obstacles [][]int) {
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		for x, loc := range line {
			if loc == '.' {
				continue
			}

			var pos []int
			pos = append(pos, x)
			pos = append(pos, y)

			if loc == '#' {
				obstacles = append(obstacles, pos)
			} else if loc == '^' {
				start = pos
			}
		}
	}

	return start, obstacles
}
