package day5

import (
	"fmt"
	"testing"
)

func TestWinds(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	matrix := ParseWindMatrix(input, false)

	thresholdCount := CountTresholdValues(matrix, 2)

	if thresholdCount != 5 {
		t.Error("UnexpectedThreshold Count for 2! Got", thresholdCount, "expected", 5)
	}
}

func TestWindsDiagonal(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	matrix := ParseWindMatrix(input, true)

	thresholdCount := CountTresholdValues(matrix, 2)

	for _, r := range matrix[0:10] {
		fmt.Println(r[0:10])
	}

	if thresholdCount != 12 {
		t.Error("UnexpectedThreshold Count for 2! Got", thresholdCount, "expected", 5)
	}
}
