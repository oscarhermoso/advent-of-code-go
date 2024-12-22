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

func part1(input string) int {
	terrain, starts, finishes := parseInput(input)
	ans := 0

	dimensions := image.Rect(0, 0, len(terrain[0]), len(terrain))

	dirs := []image.Point{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	var dfs func(height int, loc image.Point)

	dfs = func(height int, loc image.Point) {
		if !loc.In(dimensions) ||
			terrain[loc.Y][loc.X] != height {
			return
		}
		if height == 9 {
			if !finishes[loc] {
				finishes[loc] = true
				ans += 1
				return
			}
			return
		}
		for _, dir := range dirs {
			dfs(height+1, loc.Add(dir))
		}
	}

	for _, pos := range starts {
		dfs(0, pos)
		for loc := range finishes {
			finishes[loc] = false
		}
	}

	return ans
}

func part2(input string) int {
	terrain, starts, finishes := parseInput(input)
	ans := 0

	dimensions := image.Rect(0, 0, len(terrain[0]), len(terrain))

	dirs := []image.Point{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	var dfs func(height int, loc image.Point)

	dfs = func(height int, loc image.Point) {
		if !loc.In(dimensions) ||
			terrain[loc.Y][loc.X] != height {
			return
		}
		if height == 9 {
			ans += 1
			// if !finishes[loc] {
			// 	finishes[loc] = true
			// 	return
			// }
			return
		}
		for _, dir := range dirs {
			dfs(height+1, loc.Add(dir))
		}
	}

	for _, pos := range starts {
		dfs(0, pos)
		for loc := range finishes {
			finishes[loc] = false
		}
	}

	return ans
}

func parseInput(input string) (terrain [][]int, starts []image.Point, finishes map[image.Point]bool) {
	finishes = make(map[image.Point]bool)

	for y, line := range strings.Split(input, "\n") {
		var heights []int
		for x, char := range line {
			height, _ := strconv.Atoi(string(char))
			heights = append(heights, height)
			if height == 0 {
				starts = append(starts, image.Point{x, y})
			} else if height == 9 {
				finish := image.Point{x, y}
				finishes[finish] = false
			}
		}
		terrain = append(terrain, heights)
	}
	return
}
