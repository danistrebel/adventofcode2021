package day6

import "testing"

func TestLanternFishPopulation(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	population := SimulateLatnernFishPopulation(input, 18)
	expectedPopulationSize := 26
	if len(population) != expectedPopulationSize {
		t.Error("invalid simulation population size. Got", len(population), "expected", expectedPopulationSize)
	}
}
