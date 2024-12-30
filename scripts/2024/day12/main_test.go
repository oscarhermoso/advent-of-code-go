package main

import (
	"testing"
)

var example = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  1930,
		},
		{
			name:  "actual",
			input: input,
			want:  1477762,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "example ABCD",
			input: `AAAA
BBCD
BBCC
EEEC`,
			want: 80,
		},
		{
			name: "example OXO",
			input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			want: 436,
		},
		{
			name: "example E",
			input: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			want: 236,
		},
		{
			name: "example ABBA",
			input: `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
			want: 368,
		},
		{
			name: "example",
			input: `AABB
AABB
AACC
AACC`,
			want: 1206,
		},
		{
			name:  "example",
			input: example,
			want:  1206,
		},
		{
			name:  "actual",
			input: input,
			want:  0,
			// too low 913094
			// too low 919189
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
