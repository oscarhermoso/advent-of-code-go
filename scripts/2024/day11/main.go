package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	ans := 0
	stones := strings.Fields(input)
	blinks := 25

	var cache []map[string]int
	for i := 0; i < blinks; i++ {
		cache = append(cache, make(map[string]int))
	}
	for _, stone := range stones {
		ans += dfs(stone, cache, blinks)
	}
	return ans
}

func part2(input string) int {
	ans := 0
	stones := strings.Fields(input)
	blinks := 75

	var cache []map[string]int
	for i := 0; i < blinks; i++ {
		cache = append(cache, make(map[string]int))
	}
	for _, stone := range stones {
		ans += dfs(stone, cache, blinks)
	}
	return ans
}

func dfs(stone string, cache []map[string]int, i int) (ans int) {
	if i == 0 {
		return 1
	}
	i -= 1
	if cached, ok := cache[i][stone]; ok {
		return cached
	}
	if stone == "0" {
		ans = dfs("1", cache, i)
		cache[i][stone] = ans
		return ans
	} else if len(stone)%2 == 0 {
		l := stone[:len(stone)/2]
		r, _ := strconv.Atoi(stone[len(stone)/2:])
		ans += dfs(l, cache, i)
		ans += dfs(fmt.Sprint(r), cache, i)
		cache[i][stone] = ans
		return ans
	} else {
		mult, _ := strconv.ParseInt(stone, 10, 64)
		ans = dfs(fmt.Sprint(mult*2024), cache, i)
		cache[i][stone] = ans
		return ans
	}
}
