package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
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

	switch part {
	case 1:
		ans := part1(input)
		fmt.Println("Output:", ans)
	case 2:
		ans := part2(input)
		fmt.Println("Output:", ans)
	default:
		log.Fatalf("invalid part: %d", part)
	}
}

type MathProblem struct {
	Numbers  []int64
	Operator rune
}

func parseInput(input string) (ans []MathProblem) {
	for i, line := range strings.Split(input, "\n") {
		fields := strings.Fields(line)

		for j, str := range fields {
			var prob MathProblem

			// first line creates entries
			if i == 0 {
				prob = MathProblem{}
				ans = append(ans, prob)
			} else {
				prob = ans[j]
			}

			if strings.ContainsAny(str, "0123456789") {
				num, err := strconv.ParseInt(str, 10, 64)
				if err != nil {
					log.Fatalf("parseInt error: %v", err)
				}
				prob.Numbers = append(prob.Numbers, num)
			} else if strings.Contains(str, "+") {
				prob.Operator = '+'
			} else if strings.Contains(str, "*") {
				prob.Operator = '*'
			}

			// store updated copy
			ans[j] = prob
		}
	}
	return ans
}

func part1(input string) int64 {
	probs := parseInput(input)
	var total int64

	for _, prob := range probs {
		var wip int64

		if prob.Operator == '*' {
			wip = 1
			for _, n := range prob.Numbers {
				wip *= n
			}
		} else {
			for _, n := range prob.Numbers {
				wip += n
			}
		}

		total += wip
	}
	return total
}

func part2(input string) int64 {
	parsed := parseInput(input)
	_ = parsed

	// TODO: implement AoC part 2 logic
	return 0
}
