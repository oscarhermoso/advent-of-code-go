package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image"
	"regexp"
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
	challenges := parseInput(input)

	ans := 0

	for _, c := range challenges {
		// Solve linear equations for a, b:
		//   a * c.A.X + b * c.B.X = c.Prize.X
		//   a * c.A.Y + b * c.B.Y = c.Prize.Y
		//
		// Solve for b in in terms of a:
		//   b = (c.Prize.X - a * c.A.X) / c.B.X
		//   b = (c.Prize.Y - a * c.A.Y) / c.B.Y
		//
		// LHS = RHS:
		//   (c.Prize.X - a * c.A.X) / c.B.X = (c.Prize.Y - a * c.A.Y) / c.B.Y
		//
		// Solve for a in terms of constant factors:
		//   (c.Prize.X - a * c.A.X) * c.B.Y  = (c.Prize.Y - a * c.A.Y) * c.B.X
		//   c.Prize.X * c.B.Y - a * c.A.X * c.B.Y = c.Prize.Y * c.B.X - a * c.A.Y * c.B.X
		//   a * c.A.X * c.B.Y - a * c.A.Y * c.B.X = c.Prize.X * c.B.Y - c.Prize.Y * c.B.X
		//   a * (c.A.X * c.B.Y - c.A.Y * c.B.X) = c.Prize.X * c.B.Y - c.Prize.Y * c.B.X
		//   a := (c.Prize.X*c.B.Y - c.Prize.Y*c.B.X) / (c.A.X*c.B.Y - c.A.Y*c.B.X)
		a := (c.Prize.X*c.B.Y - c.Prize.Y*c.B.X) / (c.A.X*c.B.Y - c.A.Y*c.B.X)

		if c.B.X == 0 {
			continue
		}
		// Back-substitute b
		b := (c.Prize.X - a*c.A.X) / c.B.X

		if (a >= 0) && (b >= 0) && c.A.Mul(a).Add(c.B.Mul(b)) == c.Prize {
			ans += 3*a + b
		}
	}

	return ans
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

type Challenge struct {
	A     image.Point
	B     image.Point
	Prize image.Point
}

func parseInput(input string) (challenges []Challenge) {
	re := regexp.MustCompile(`Button A: X\+([0-9]+), Y\+([0-9]+)
Button B: X\+([0-9]+), Y\+([0-9]+)
Prize: X=([0-9]+), Y=([0-9]+)
`)

	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		ax, _ := strconv.Atoi(match[1])
		ay, _ := strconv.Atoi(match[2])
		a := image.Point{ax, ay}
		bx, _ := strconv.Atoi(match[3])
		by, _ := strconv.Atoi(match[4])
		b := image.Point{bx, by}
		px, _ := strconv.Atoi(match[5])
		py, _ := strconv.Atoi(match[6])
		p := image.Point{px, py}
		challenges = append(challenges, Challenge{a, b, p})
	}

	return challenges
}
