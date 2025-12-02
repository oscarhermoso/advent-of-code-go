package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image"
	"log"
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

type Robot struct {
	position image.Point
	velocity image.Point
}

func parseInput(input string) (ans []Robot) {
	for _, line := range strings.Split(input, "\n") {
		var px, py, vx, vy int
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

		if err != nil {
			log.Fatalf("failed to parse input: %v", err)
		}
		// log.Printf("Parsed robot { pos: %d, %d; vel: %d, %d}", px, py, vx, vy)
		ans = append(ans, Robot{
			position: image.Pt(px, py),
			velocity: image.Pt(vx, vy),
		})
	}
	return ans
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed

	// example
	width := 11
	height := 7

	// width := 101
	// height := 103
	var q1, q2, q3, q4 int

	for _, robot := range parsed {
		robot.position = robot.position.Add(robot.velocity.Mul(100)).Mod(image.Rect(0, 0, width, height))

		fmt.Printf("Robot position: %s")

		if robot.position.X < width/2 {
			if robot.position.Y < height/2 {
				q1 += 1
			} else if robot.position.Y > height/2+1 {
				q2 += 1
			}
		} else if robot.position.X > width/2+1 {
			if robot.position.Y < height/2 {
				q3 += 1
			} else if robot.position.Y > height/2+1 {
				q4 += 1
			}
		}
	}

	return q1 * q2 * q3 * q4 // too low 201437250
}

func part2(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}
