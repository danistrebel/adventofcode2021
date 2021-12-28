package day7

import "testing"

func TestOptimalCrabsPosition(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	fuel := OptimalCrabsPosition(input)

	expectedFuel := 37

	if fuel != expectedFuel {
		t.Error("Wrong Fuel Calculated. Got", fuel, "expected", expectedFuel)
	}

}

func TestOptimalCrabsPositionTwo(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	fuel := OptimalCrabsPositionTwo(input)

	expectedFuel := 168

	if fuel != expectedFuel {
		t.Error("Wrong Fuel Calculated. Got", fuel, "expected", expectedFuel)
	}

}
