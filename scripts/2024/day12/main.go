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
		fmt.Printf("%v\n", region.sides)
	}
	return price
}

type Region struct {
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
			region := Region{[]image.Point{pos}, 0, 0, []image.Rectangle{}}
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
			region.sides = mergeSides(region.sides)
			regions = append(regions, region)
		}
	}

	return regions
}

func mergeSides(sides []image.Rectangle) (mergedSides []image.Rectangle) {
	if len(sides) < 2 {
		return sides
	}
	if len(sides) >= 2 {
		mergedSides = append(mergedSides, mergeSides(sides[:len(sides)/2])...)
		mergedSides = append(mergedSides, mergeSides(sides[len(sides)/2:])...)
	}

	remaining := true
mergeLoop:
	for remaining {
		for i, side1 := range sides {
			for j, side2 := range sides {
				if i == j {
					continue
				}

				if (side1.Min.Eq(side2.Max) && (side1.Max.X == side2.Min.X || side1.Max.Y == side2.Min.Y)) ||
					(side1.Min.Eq(side2.Min) && (side1.Max.X == side2.Max.X || side1.Max.Y == side2.Max.Y)) ||
					(side1.Max.Eq(side2.Max) && (side1.Min.X == side2.Min.X || side1.Min.Y == side2.Min.Y)) {

					fmt.Printf("mergerino %s and %s\n", side1, side2)

					if side1.Min.X > side2.Min.X {
						side1.Min.X = side2.Min.X
					}
					if side1.Min.Y > side2.Min.Y {
						side1.Min.Y = side2.Min.Y
					}
					if side1.Max.X < side2.Max.X {
						side1.Max.X = side2.Max.X
					}
					if side1.Max.Y < side2.Max.Y {
						side1.Max.Y = side2.Max.Y
					}

					fmt.Printf("merged into %s\n", side1)

					sides[i] = side1
					mergedSides = append(mergedSides, sides[:j]...)
					mergedSides = append(mergedSides, sides[j+1:]...)

					fmt.Printf("sides are %v\n", mergedSides)
					sides = mergedSides
					mergedSides = []image.Rectangle{}
					continue mergeLoop
				}
			}
		}
		remaining = false
	}
	// fmt.Printf("done mergin sides\n")
	return sides
}
