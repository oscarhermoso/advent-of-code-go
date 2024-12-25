package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	connections := parseInput(input)
	found := make(map[string]struct{})

	for c0, adj0 := range connections {
		if c0[0] != 't' {
			continue
		}
		for c1 := range adj0 {
			if c1 == c0 {
				continue
			}
			adj1 := connections[c1]
			for c2 := range adj1 {
				if c2 == c1 || c2 == c0 {
					continue
				}
				if _, ok := adj0[c2]; !ok {
					continue
				}
				f := []string{c0, c1, c2}
				slices.Sort(f)
				found[fmt.Sprintf("%v", f)] = struct{}{}
			}
		}
	}
	// fmt.Printf("%v", found)
	return len(found)
}

func part2(input string) int {
	connections := parseInput(input)
	_ = connections

	return 0
}

func parseInput(input string) (connections map[string]map[string]struct{}) {
	connections = make(map[string]map[string]struct{})

	for _, line := range strings.Split(input, "\n") {
		computers := strings.Split(line, "-")
		c0 := computers[0]
		c1 := computers[1]

		if adj, ok := connections[c0]; ok {
			adj[c1] = struct{}{}
		} else {
			connections[c0] = map[string]struct{}{c1: {}}
		}
		if adj, ok := connections[c1]; ok {
			adj[c0] = struct{}{}
		} else {
			connections[c1] = map[string]struct{}{c0: {}}
		}
	}

	return connections
}
