package day6

import "testing"

func TestLanternFishPopulation(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	population := SimulateLatnernFishPopulation(input, 18)
	expectedPopulationSize := 26
	if population != expectedPopulationSize {
		t.Error("invalid simulation population size after 18 days. Got", population, "expected", expectedPopulationSize)
	}

}

func TestCompactLanternFishPopulation(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	population := CompactSimulateLatnernFishPopulation(input, 18)
	expectedPopulationSize := uint64(26)
	if population != expectedPopulationSize {
		t.Error("invalid simulation population size after 18 days. Got", population, "expected", expectedPopulationSize)
	}

	population256 := CompactSimulateLatnernFishPopulation(input, 256)
	expectedPopulationSize256 := uint64(26984457539)
	if population256 != expectedPopulationSize256 {
		t.Error("invalid simulation population size after 256 days. Got", population256, "expected", expectedPopulationSize256)
	}

}
