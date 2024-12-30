package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image"
	"slices"
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
	regions := parseInput(input)
	price := 0
	for _, region := range regions {
		price += region.area * region.perimeter
	}
	return price
}

func part2(input string) int {
	regions := parseInput(input)
	price := 0
	for _, region := range regions {
		price += region.area * len(region.sides)
		// fmt.Printf("%c (%d): %v\n", region.char, len(region.sides), region.sides)
	}

	var debug [][]rune

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		debug = append(debug, []rune(strings.Repeat(" ", len(line)*2+1)))
		debugLine := make([]rune, len(lines[0])*2+1)
		for i, char := range line {
			debugLine[i*2] = ' '
			debugLine[i*2+1] = char
		}
		debugLine[len(line)*2] = ' '
		debug = append(debug, debugLine)
	}
	debug = append(debug, []rune(strings.Repeat(" ", len(debug[0]))))

	// for _, reg := range regions {
	// 	for _, side := range reg.sides {
	// 		// for x := side.Min.X * 2; x < side.Max.X*2; x++ {
	// 		// 	debug[side.Min.Y][x] = '-'
	// 		// }
	// 		// for y := side.Min.Y * 2; y < side.Max.Y*2; y++ {
	// 		// 	debug[y][side.Min.X] = '|'
	// 		// }

	// 		debug[side.Min.Y*2][side.Min.X*2] = '+'
	// 		debug[side.Max.Y*2][side.Max.X*2] = '+'
	// 	}
	// }

	for _, debugLine := range debug {
		fmt.Printf("%s (%d)\n", string(debugLine), len(debugLine))
	}

	return price
}

var regionId int = 0

type Region struct {
	id        int
	char      rune
	points    []image.Point
	perimeter int
	area      int
	sides     []image.Rectangle
}

func parseInput(input string) (regions []Region) {
	lines := strings.Split(input, "\n")
	points := make(map[image.Point]Region)

	dirs := []image.Point{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}
	sideDirs := []image.Rectangle{
		image.Rect(0, 1, 1, 1), // bottom
		image.Rect(1, 0, 1, 1), // right
		image.Rect(0, 0, 1, 0), // top
		image.Rect(0, 0, 0, 1), // left
	}

	rect := image.Rect(0, 0, len(lines[0]), len(lines))

	for y, line := range lines {
		for x, char := range line {
			pos := image.Pt(x, y)
			if _, ok := points[pos]; ok {
				continue
			}
			region := Region{regionId, char, []image.Point{pos}, 0, 0, []image.Rectangle{}}
			regionId++
			for i := 0; i < len(region.points); i++ {
				pos := region.points[i]
				region.area += 1
				points[pos] = region
				for j, dir := range dirs {
					neighbour := pos.Add(dir)
					if slices.Contains(region.points, neighbour) {
						continue
					}
					if neighbour.In(rect) && rune(lines[neighbour.Y][neighbour.X]) == char {
						region.points = append(region.points, neighbour)
					} else {
						region.perimeter += 1
						region.sides = append(region.sides, sideDirs[j].Add(pos))
					}
				}
			}
			mobiusPosts := make(map[image.Point]struct{})
			for _, pos := range region.points {
				if post := pos.Add(image.Pt(1, 1)); slices.Contains(region.points, post) &&
					(!slices.Contains(region.points, pos.Add(image.Pt(0, 1))) && //    A|B
						!slices.Contains(region.points, pos.Add(image.Pt(1, 0)))) { // C|A
					mobiusPosts[post] = struct{}{}
					// fmt.Printf("found mobius fence: %s", post)
				} else if post := pos.Add(image.Pt(-1, 1)); slices.Contains(region.points, post) &&
					(!slices.Contains(region.points, pos.Add(image.Pt(-1, 0))) && //   B|A
						!slices.Contains(region.points, pos.Add(image.Pt(0, 1)))) { // A|C
					mobiusPosts[post] = struct{}{}
					// fmt.Printf("found mobius fence: %s", post)
				}
			}
			region.sides = mergeSides(region.sides, mobiusPosts)
			regions = append(regions, region)
		}
	}

	return regions
}

func mergeSides(sides []image.Rectangle, mobiusPosts map[image.Point]struct{}) (mergedSides []image.Rectangle) {
	if len(sides) < 2 {
		return sides
	}
	if len(sides) >= 2 {
		mergedSides = append(mergedSides, mergeSides(sides[:len(sides)/2], mobiusPosts)...)
		mergedSides = append(mergedSides, mergeSides(sides[len(sides)/2:], mobiusPosts)...)
	}
	remaining := true
mergeLoop:
	for remaining {
		remaining = false
		for i, side1 := range mergedSides {
			for j, side2 := range mergedSides {
				if i == j {
					continue
				}
				if side1.Max.Eq(side2.Min) && (side1.Min.X == side2.Max.X || side1.Min.Y == side2.Max.Y) ||
					side1.Max.Eq(side2.Max) && (side1.Min.X == side2.Min.X || side1.Min.Y == side2.Min.Y) ||
					side1.Min.Eq(side2.Min) && (side1.Max.X == side2.Max.X || side1.Max.Y == side2.Max.Y) {
					// guard mobius fence
					if _, ok := mobiusPosts[side1.Max]; ok {
						continue
					}

					// merge
					mergedSides[i] = image.Rect(side1.Min.X, side1.Min.Y, side2.Max.X, side2.Max.Y)
					mergedSides[j] = mergedSides[len(mergedSides)-1]
					mergedSides = mergedSides[:len(mergedSides)-1]
					remaining = true
					continue mergeLoop
				}
			}
		}
	}
	return mergedSides
}
