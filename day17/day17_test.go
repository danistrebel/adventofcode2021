package day17

import "testing"

func TestTrajectory(t *testing.T) {
	s := speed{7, 2}
	expectedTrajectory := []coord{
		{7, 2},
		{13, 3},
		{18, 3},
		{22, 2},
	}

	position := coord{0, 0}
	for i := 0; i < len(expectedTrajectory); i++ {
		position, s = simulateNext(position, s)
		if position != expectedTrajectory[i] {
			t.Error("Unexpected Position at step", i, "got", position, "got", expectedTrajectory[i])
		}
	}
}

func TestAngle(t *testing.T) {
	s := speed{7, 2}
	tar := target{20, 30, -10, -5}
	hit := testAngle(s, tar)
	if !hit {
		t.Error("Expected target hit but got miss")
	}
}

func TestAngleHigh(t *testing.T) {
	s := speed{6, 9}
	tar := target{20, 30, -10, -5}
	hit := testAngle(s, tar)

	if !hit {
		t.Error("Expected target miss but got hit")
	}
}

func TestUndershoot(t *testing.T) {
	s := speed{6, 9}
	tar := target{20, 30, -10, -5}
	hit := testAngle(s, tar)

	if hit {
		t.Error("Expected target miss but got hit")
	}
}

func TestOvershoot(t *testing.T) {
	s := speed{17, -4}
	tar := target{20, 30, -10, -5}
	hit := testAngle(s, tar)

	if hit {
		t.Error("Expected target miss but got hit")
	}
}

func TestHighestY(t *testing.T) {
	expectedMaxY := 45
	tar := target{20, 30, -10, -5}
	maxY := findHighestY(tar)

	if maxY != expectedMaxY {
		t.Error("Unexpected max Y. Got", maxY, "expected", expectedMaxY)
	}
}
