package util_test

import (
	"reflect"
	"testing"

	"github.com/lewinski/advent-of-code/util"
)

func TestIntFields(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{"", []int{}},
		{"0", []int{0}},
		{"1 2 3", []int{1, 2, 3}},
		{"1     3", []int{1, 3}},
		{"foo 3 bar", []int{0, 3, 0}},
	}

	for _, tt := range tests {
		got := util.IntFields(tt.input)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("IntFields(%s): expected %v got %v", tt.input, tt.want, got)
		}
	}
}
