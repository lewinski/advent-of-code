package util_test

import (
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestFirstRune(t *testing.T) {
	tests := []struct {
		input string
		want  rune
	}{
		{"", 0},
		{"abc", 'a'},
		{"🐻🐷", '🐻'},
	}

	for _, tt := range tests {
		got := util.FirstRune(tt.input)
		if got != tt.want {
			t.Errorf("FirstRune(%s): expected '%c' got '%c'", tt.input, tt.want, got)
		}
	}
}

func TestLastRune(t *testing.T) {
	tests := []struct {
		input string
		want  rune
	}{
		{"", 0},
		{"abc", 'c'},
		{"🐻🐷", '🐷'},
	}

	for _, tt := range tests {
		got := util.LastRune(tt.input)
		if got != tt.want {
			t.Errorf("LastRune(%s): expected '%c' got '%c'", tt.input, tt.want, got)
		}
	}
}
