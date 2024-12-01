package main

import "github.com/oscarhermoso/advent-of-code-go/scripts/aoc"

func main() {
	day, year, cookie := aoc.ParseFlags()
	aoc.GetInput(day, year, cookie)
}
