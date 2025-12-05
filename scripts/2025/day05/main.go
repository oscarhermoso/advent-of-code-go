package main

import (
	"cmp"
	_ "embed"
	"flag"
	"fmt"
	"math"
	"slices"
	"strconv"
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

func parseInput(input string) (ranges [][]int64, ingredients []int64) {
	for _, line := range strings.Split(input, "\n") {
		if strings.ContainsRune(line, '-') {
			ingredientIds := strings.Split(line, "-")
			lowerId, _ := strconv.ParseInt(ingredientIds[0], 10, 64)
			upperId, _ := strconv.ParseInt(ingredientIds[1], 10, 64)
			ranges = append(ranges, []int64{lowerId, upperId})
		} else if strings.ContainsAny(line, "1234567890") {
			ingredient, _ := strconv.ParseInt(line, 10, 64)
			ingredients = append(ingredients, ingredient)
		}
	}
	return ranges, ingredients
}

func part1(input string) int {
	ranges, ingredients := parseInput(input)
	var fresh int

	for _, i := range ingredients {
		for _, r := range ranges {
			if i >= r[0] && i <= r[1] {
				fresh += 1
				break
			}
		}
	}

	return fresh
}

func part2(input string) int64 {
	ranges, _ := parseInput(input)

	slices.SortFunc(ranges, func(a, b []int64) int {
		return cmp.Compare(a[0], b[0])
	})

	var history [][]int64
	for _, r := range ranges {
		history = append(history, r)
		for {
			matched := false
			for i, h1 := range history {
				for j, h2 := range history {
					if h1[1] >= h2[0] && h1[0] <= h2[1] {
						if i >= j {
							continue
						}
						newMin := int64(math.Min(float64(h1[0]), float64(h2[0])))
						newMax := int64(math.Max(float64(h1[1]), float64(h2[1])))

						newHistory := slices.Clone(history[:i])
						if i+1 != j {
							newHistory = append(newHistory, history[i+1:j]...)
						}
						newHistory = append(newHistory, []int64{newMin, newMax})
						if j < len(history)-1 {
							newHistory = append(newHistory, history[j+1:]...)
						}
						history = newHistory
						matched = true
					}
				}
			}
			if !matched {
				break
			}
		}
	}

	var fresh int64
	for _, h := range history {
		// fmt.Printf("fresh %d-%d\n", h[0], h[1])
		fresh += h[1] - h[0] + 1
	}

	return fresh
}
