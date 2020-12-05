package main

import (
	"testing"
)

func TestDecodeSeat(t *testing.T) {
	tests := []struct {
		pass     string
		row, col int
		id       int
	}{
		{"FBFBBFFRLR", 44, 5, 357},
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}

	for i, tt := range tests {
		row, col := decodeSeat(tt.pass)
		if row != tt.row {
			t.Errorf("test %d: row incorrect; want %d, got %d", i, tt.row, row)
		}
		if col != tt.col {
			t.Errorf("test %d: col incorrect; want %d, got %d", i, tt.col, col)
		}

		id := seatID(row, col)
		if id != tt.id {
			t.Errorf("test %d: id incorrect; want %d, got %d", i, tt.id, id)
		}
	}
}
