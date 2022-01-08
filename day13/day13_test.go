package day13

import (
	"testing"
)

func TestSingleFold(t *testing.T) {
	inputLines := []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
		"",
		"fold along y=7",
		"fold along x=5",
	}

	grid, folds := parseInput(inputLines)

	gridFold := foldGrid(grid, folds[0])

	count := countDots(gridFold)
	expectedCount := 17

	if count != expectedCount {
		t.Error("Unexpected Count after first fold. Got", count, "expected", expectedCount)
	}
}
