package main

import (
	"cmp"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math"
	"slices"
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

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

type Pair struct {
	a, b     *[]int
	distance float64
}

func parseInput(input string) (ans [][]int) {
	for _, line := range strings.Split(input, "\n") {
		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			log.Fatalf("failed to parse input: %v", err)
		}
		ans = append(ans, []int{x, y, z})
	}
	return ans
}

func part1(input string) int {
	parsed := parseInput(input)
	var pairs []Pair

	for i, a := range parsed[:len(parsed)-1] {
		for _, b := range parsed[i+1:] {
			pairs = append(pairs, Pair{&a, &b, Distance3D(a, b)})
		}
	}

	slices.SortFunc(pairs, func(a Pair, b Pair) int {
		return cmp.Compare(a.distance, b.distance)
	})

	circuits := make(map[string][]string)

	joined := 0
	for _, pair := range pairs {
		a := fmt.Sprintf("%v", *pair.a)
		b := fmt.Sprintf("%v", *pair.b)
		circuitA, okA := circuits[a]
		circuitB, okB := circuits[b]

		if okA && okB && !slices.Contains(circuitA, b) && !slices.Contains(circuitB, a) {
			newCircuit := append(circuitA, circuitB...)
			newCircuit = slices.Compact(newCircuit)
			for _, coord := range newCircuit {
				circuits[coord] = newCircuit
			}
		} else if !okA && !okB {
			newCircuit := []string{a, b}
			circuits[a] = newCircuit
			circuits[b] = newCircuit
		} else if !okA && !slices.Contains(circuitB, a) {
			newCircuit := append(circuitB, a)
			for _, coord := range newCircuit {
				circuits[coord] = newCircuit
			}
		} else if !okB && !slices.Contains(circuitA, b) {
			newCircuit := append(circuitA, b)
			for _, coord := range newCircuit {
				circuits[coord] = newCircuit
			}
		}
		joined += 1
		if joined >= 1000 {
			break
		}
	}

	visited := make(map[string]bool)
	longestCircuits := make([]int, 0, len(circuits))
	for _, c := range circuits {
		str := fmt.Sprintf("%v", c)
		_, ok := visited[str]
		if !ok {
			visited[str] = true
			longestCircuits = append(longestCircuits, len(c))
			// fmt.Println(c)
		}
	}
	slices.SortFunc(longestCircuits, func(a int, b int) int {
		return -cmp.Compare(a, b)
	})
	fmt.Printf("lengths of curcuits: %d, %d, %d, %d\n", longestCircuits[0], longestCircuits[1], longestCircuits[2], longestCircuits[3])

	return longestCircuits[0] * longestCircuits[1] * longestCircuits[2]
}

func part2(input string) int {
	parsed := parseInput(input)
	var pairs []Pair

	for i, a := range parsed[:len(parsed)-1] {
		for _, b := range parsed[i+1:] {
			pairs = append(pairs, Pair{&a, &b, Distance3D(a, b)})
		}
	}

	slices.SortFunc(pairs, func(a Pair, b Pair) int {
		return cmp.Compare(a.distance, b.distance)
	})

	circuits := make(map[string][]string)

	for _, pair := range pairs {
		a := fmt.Sprintf("%v", *pair.a)
		b := fmt.Sprintf("%v", *pair.b)
		circuitA, okA := circuits[a]
		circuitB, okB := circuits[b]
		var newCircuit []string

		if okA && okB && !slices.Contains(circuitA, b) && !slices.Contains(circuitB, a) {
			newCircuit = append(circuitA, circuitB...)
			newCircuit = slices.Compact(newCircuit)
		} else if !okA && !okB {
			newCircuit = []string{a, b}
		} else if !okA && !slices.Contains(circuitB, a) {
			newCircuit = append(circuitB, a)
		} else if !okB && !slices.Contains(circuitA, b) {
			newCircuit = append(circuitA, b)
		}

		for _, coord := range newCircuit {
			circuits[coord] = newCircuit
		}

		if len(newCircuit) == len(parsed) {
			return (*pair.a)[0] * (*pair.b)[0]
		}
	}

	return -1
}

func Distance3D(a []int, b []int) float64 {
	dx := float64(b[0] - a[0])
	dy := float64(b[1] - a[1])
	dz := float64(b[2] - a[2])
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
