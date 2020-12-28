package main

import (
	"testing"

	"github.com/alexchao26/advent-of-code-go/util"
)

func Test_part1(t *testing.T) {
	type args struct {
		input  string
		myHP   int
		myMana int
		part   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{"Hit Points: 13\nDamage: 8", 10, 250, 1},
			want: spellsMap["Poison"].cost + spellsMap["Magic Missile"].cost,
		},
		{"actual", args{util.ReadFile("input.txt"), 50, 500, 1}, 953},
		{"actual", args{util.ReadFile("input.txt"), 50, 500, 2}, 1289},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input, tt.args.myHP, tt.args.myMana, tt.args.part); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
