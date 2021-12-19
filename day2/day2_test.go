package day2

import "testing"

func TestCoordinates(t *testing.T) {
	x, d := calculateCoordinates([]string{
		"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2",
	})
	if x != 15 {
		t.Error("invalid horizontal position! Expected 15 but got: ", x)
	}
	if d != 10 {
		t.Error("invalid depth! Expected 10 but got: ", d)
	}
}

func TestAimCoordinates(t *testing.T) {
	x, d := calculateAimCoordinates([]string{
		"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2",
	})
	if x != 15 {
		t.Error("invalid horizontal position! Expected 15 but got: ", x)
	}
	if d != 60 {
		t.Error("invalid depth! Expected 60 but got: ", d)
	}
}
