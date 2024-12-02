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
	reports := parseInput(input)

	unsafeCount := 0

eachReport:
	for _, report := range reports {
		if report[0] == report[1] {
			unsafeCount += 1
			continue
		}

		if report[0] < report[1] {
			for i, level := range report {
				if i == len(report)-1 {
					continue eachReport
				}

				if report[i+1] <= level || report[i+1] > level+3 {
					unsafeCount += 1
					continue eachReport
				}
			}
		}

		if report[0] > report[1] {
			for i, level := range report {
				if i == len(report)-1 {
					continue eachReport
				}

				if report[i+1] >= level || report[i+1] < level-3 {
					unsafeCount += 1
					continue eachReport
				}
			}
		}
	}

	return len(reports) - unsafeCount
}

func part2(input string) int {
	reports := parseInput(input)

	unsafeCount := 0

eachReport:
	for _, report := range reports {

		if isReportSafe(report) {
			continue eachReport
		}

		for i := range report {
			var cutDownReport []int
			cutDownReport = append(cutDownReport, report[:i]...)
			cutDownReport = append(cutDownReport, report[i+1:]...)

			if isReportSafe(cutDownReport) {
				continue eachReport
			}
		}

		// fmt.Printf(" (unsafe report was %#v)\n", report)
		unsafeCount += 1
	}

	return len(reports) - unsafeCount
}

func parseInput(input string) (inputs [][]int) {
	for _, line := range strings.Split(input, "\n") {
		var fields []int
		for _, field := range strings.Fields(line) {
			parsedField, _ := strconv.Atoi(field)
			fields = append(fields, parsedField)
		}
		inputs = append(inputs, fields)
	}

	return inputs
}

func isReportSafe(report []int) bool {
	if report[0] == report[1] {
		return false
	}

	if report[0] < report[1] {
		for i, level := range report {
			if i == len(report)-1 {
				return true
			}

			if report[i+1] <= level || report[i+1] > level+3 {
				// fmt.Printf("unsafe because of %d %d\n", level, report[i+1])
				return false
			}
		}
	}

	if report[0] > report[1] {
		for i, level := range report {
			if i == len(report)-1 {
				return true
			}

			if report[i+1] >= level || report[i+1] < level-3 {
				// fmt.Printf("unsafe because of %d %d\n", level, report[i+1])
				return false
			}
		}
	}

	panic("unreachable")
}
