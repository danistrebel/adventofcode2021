package day9

import "testing"

func TestTotalRiskLevel(t *testing.T) {
	input := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	riskLevel := totalRiskLevel(input)

	expectedRiskLevel := 15

	if riskLevel != expectedRiskLevel {
		t.Error("Unexpected Risk Level. Got", riskLevel, "expected", expectedRiskLevel)
	}
}

func TestBaisinSize(t *testing.T) {
	input := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	size := largestBaisins(input)

	expectedSize := 1134

	if size != expectedSize {
		t.Error("Unexpected Total Baisin Size. Got", size, "expected", expectedSize)
	}
}
