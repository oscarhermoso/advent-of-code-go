package main

import (
	"testing"
)

var example = `1
10
100
2024`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{
			name:  "example",
			input: example,
			want:  37327623,
		},
		{
			name:  "actual",
			input: input,
			want:  17163502021,
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
		want  int64
	}{
		{
			name: "example",
			input: `1
2
3
2024`,
			want: 23,
		},
		{
			name:  "actual",
			input: input,
			want:  1938,
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
