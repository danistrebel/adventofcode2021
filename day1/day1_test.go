package day1

import "testing"

func TestSonar(t *testing.T) {
	testLog := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	calculated := countIncreases(testLog)
	if calculated != 7 {
		t.Errorf("invalid increase count. Got: %d, expected: 7", calculated)
	}
}

func TestSonarSlidingWindow(t *testing.T) {
	testLog := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	calculated := countWindowIncreases(testLog, 3)
	if calculated != 5 {
		t.Errorf("invalid increase count. Got: %d, expected: 5", calculated)
	}
}
