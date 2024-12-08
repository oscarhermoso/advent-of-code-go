package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image"
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
	antennas, locations, width, height := parseInput(input)
	// fmt.Printf("dimensions: %d, %d\n", width, height)

	count := 0

	// var debug []image.Point

	for _, locs := range antennas {
		for i := 0; i < len(locs)-1; i++ {
			current := locs[i]
			remaining := locs[i+1:]

			for _, peer := range remaining {
				antinode1 := current.Add(current).Sub(peer)
				antinode2 := peer.Add(peer).Sub(current)

				if antinode1.In(image.Rect(0, 0, width, height)) {
					antinode, found := locations[antinode1]
					if !found || !antinode {
						locations[antinode1] = ANTINODE
						count += 1
						// fmt.Printf("antinode found: %v\n", antinode1)
						// debug = append(debug, antinode2)
					}
				}

				if antinode2.In(image.Rect(0, 0, width, height)) {
					antinode, found := locations[antinode2]
					if !found || !antinode {
						locations[antinode2] = ANTINODE
						count += 1
						// fmt.Printf("antinode found: %v\n", antinode2)
						// debug = append(debug, antinode2)
					}
				}
			}
		}
	}

	// fmt.Printf("%v", debug)

	return count
}

func part2(input string) int {
	antennas, locations, width, height := parseInput(input)
	// fmt.Printf("dimensions: %d, %d\n", width, height)

	count := 0

	// var debug []image.Point

	for _, locs := range antennas {
		for i := 0; i < len(locs)-1; i++ {
			current := locs[i]
			remaining := locs[i+1:]

			for _, peer := range remaining {
				loc := current
				for {
					if loc.In(image.Rect(0, 0, width, height)) {
						antinode, found := locations[loc]
						if !found || !antinode {
							locations[loc] = ANTINODE
							count += 1
							// fmt.Printf("antinode found: %v\n", loc)
							// debug = append(debug, loc)
						}

					} else {
						break
					}
					loc = loc.Add(current).Sub(peer)
				}

				loc = peer
				for {
					if loc.In(image.Rect(0, 0, width, height)) {
						antinode, found := locations[loc]
						if !found || !antinode {
							locations[loc] = ANTINODE
							count += 1
							// fmt.Printf("antinode found: %v\n", loc)
							// debug = append(debug, loc)
						}
					} else {
						break
					}
					loc = loc.Add(peer).Sub(current)
				}
			}
		}
	}

	// fmt.Printf("%v", debug)

	return count
}

const (
	ANTENNA  = false
	ANTINODE = true
)

func parseInput(input string) (antennas map[rune][]image.Point, locations map[image.Point]bool, width, height int) {
	antennas = make(map[rune][]image.Point)
	locations = make(map[image.Point]bool)

	lines := strings.Split(input, "\n")

	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				loc := image.Point{x, y}
				antennas[char] = append(antennas[char], loc)
				locations[loc] = ANTENNA
			}
		}
	}
	height = len(lines)
	width = len(lines[0])

	return antennas, locations, width, height
}
