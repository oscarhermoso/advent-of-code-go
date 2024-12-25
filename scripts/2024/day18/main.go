package main

import (
	"container/heap"
	_ "embed"
	"flag"
	"fmt"
	"image"
	"slices"
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

type PrioritisedPoint struct {
	point    image.Point
	priority int
}
type MinHeap []PrioritisedPoint

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(PrioritisedPoint))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func part1(input string) int {
	corrupted := parseInput(input)
	start := image.Pt(0, 0)
	fin := image.Pt(6, 6) // example dimensions
	for pos := range corrupted {
		if pos.X > 6 || pos.Y > 6 {
			fin = image.Point{70, 70} // full-sized dimensions
			break
		}
	}

	rect := image.Rect(start.X, start.Y, fin.X+1, fin.Y+1)

	dirs := []image.Point{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	// A*
	h := func(pos image.Point) int {
		delta := fin.Sub(pos)
		return delta.X + delta.Y
	}

	openSet := &MinHeap{PrioritisedPoint{start, 0}}
	heap.Init(openSet)
	cameFrom := make(map[image.Point]image.Point)
	gScore := map[image.Point]int{start: 0}        // known scores
	fScore := map[image.Point]int{start: h(start)} // guessed scores

	reconstructPath := func(cameFrom map[image.Point]image.Point, current image.Point) []image.Point {
		totalPath := []image.Point{current}
		for !start.Eq(current) {
			current = cameFrom[current]
			totalPath = append(totalPath, current)
		}
		slices.Reverse(totalPath)
		return totalPath
	}

	var finalPath []image.Point
	// fmt.Printf("corrupted: %v\n", corrupted)

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(PrioritisedPoint).point
		// fmt.Printf("%v\n", reconstructPath(cameFrom, current))
		// fmt.Printf("%v\n", gScore)
		if current.Eq(fin) {
			finalPath = reconstructPath(cameFrom, current)
			break
		}
		for _, dir := range dirs {
			neighbour := current.Add(dir)
			if !neighbour.In(rect) {
				continue
			} else if _, ok := corrupted[neighbour]; ok {
				continue
			}
			tenativeGScore := gScore[current] + 1
			if g, ok := gScore[neighbour]; !ok || tenativeGScore < g {
				cameFrom[neighbour] = current
				gScore[neighbour] = tenativeGScore
				f := tenativeGScore + h(neighbour)
				fScore[neighbour] = f
				found := false
				for _, p := range *openSet {
					if p.point.Eq(neighbour) {
						found = true
						break
					}
				}
				if !found {
					heap.Push(openSet, PrioritisedPoint{neighbour, f})
				}
			}
		}
	}

	// DEBUG
	// for y := 0; y <= fin.Y; y++ {
	// 	for x := 0; x <= fin.X; x++ {
	// 		if _, ok := corrupted[image.Pt(x, y)]; ok {
	// 			print("#")
	// 		} else if slices.Contains(finalPath, image.Pt(x, y)) {
	// 			print("O")
	// 		} else {
	// 			print(".")
	// 		}
	// 	}
	// 	print("\n")
	// }

	return len(finalPath) - 1
}

func part2(input string) string {
	corruptedSlice := parseInput2(input)

	start := image.Pt(0, 0)
	fin := image.Pt(6, 6) // example dimensions
	for _, pos := range corruptedSlice {
		if pos.X > 6 || pos.Y > 6 {
			fin = image.Point{70, 70} // full-sized dimensions
			break
		}
	}

	rect := image.Rect(start.X, start.Y, fin.X+1, fin.Y+1)

	dirs := []image.Point{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0},
	}

	// A*
	h := func(pos image.Point) int {
		delta := fin.Sub(pos)
		return delta.X + delta.Y
	}

	for i, finalCorruptedPoint := range corruptedSlice {
		corrupted := make(map[image.Point]int)
		for j, c := range corruptedSlice[:i+1] {
			corrupted[c] = j
		}

		openSet := &MinHeap{PrioritisedPoint{start, 0}}
		heap.Init(openSet)
		cameFrom := make(map[image.Point]image.Point)
		gScore := map[image.Point]int{start: 0}        // known scores
		fScore := map[image.Point]int{start: h(start)} // guessed scores

		reconstructPath := func(cameFrom map[image.Point]image.Point, current image.Point) []image.Point {
			totalPath := []image.Point{current}
			for !start.Eq(current) {
				current = cameFrom[current]
				totalPath = append(totalPath, current)
			}
			slices.Reverse(totalPath)
			return totalPath
		}

		var finalPath []image.Point
		// fmt.Printf("corrupted: %v\n", corrupted)

		for openSet.Len() > 0 {
			current := heap.Pop(openSet).(PrioritisedPoint).point
			// fmt.Printf("%v\n", reconstructPath(cameFrom, current))
			// fmt.Printf("%v\n", gScore)
			if current.Eq(fin) {
				finalPath = reconstructPath(cameFrom, current)
				break
			}
			for _, dir := range dirs {
				neighbour := current.Add(dir)
				if !neighbour.In(rect) {
					continue
				} else if _, ok := corrupted[neighbour]; ok {
					continue
				}
				tenativeGScore := gScore[current] + 1
				if g, ok := gScore[neighbour]; !ok || tenativeGScore < g {
					cameFrom[neighbour] = current
					gScore[neighbour] = tenativeGScore
					f := tenativeGScore + h(neighbour)
					fScore[neighbour] = f
					found := false
					for _, p := range *openSet {
						if p.point.Eq(neighbour) {
							found = true
							break
						}
					}
					if !found {
						heap.Push(openSet, PrioritisedPoint{neighbour, f})
					}
				}
			}
		}

		if finalPath == nil {
			// DEBUG
			for y := 0; y <= fin.Y; y++ {
				for x := 0; x <= fin.X; x++ {
					if _, ok := corrupted[image.Pt(x, y)]; ok {
						print("#")
					} else if slices.Contains(finalPath, image.Pt(x, y)) {
						print("O")
					} else {
						print(".")
					}
				}
				print("\n")
			}
			return fmt.Sprintf("%d,%d", finalCorruptedPoint.X, finalCorruptedPoint.Y)
		}
	}

	return ""
}

func parseInput(input string) (ans map[image.Point]int) {
	ans = make(map[image.Point]int)
	for i, line := range strings.Split(input, "\n") {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		ans[image.Point{x, y}] = i
		// if i >= 11 { // example
		// 	break
		// }
		if i >= 1023 { // part1
			break
		}
	}

	return ans
}

func parseInput2(input string) (ans []image.Point) {
	for _, line := range strings.Split(input, "\n") {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		ans = append(ans, image.Point{x, y})
	}

	return ans
}
