package day3

import "testing"

func TestPowerConsumption(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	gamma, epsilon := powerConsumptionRates(input)

	if gamma != 22 {
		t.Errorf("Unexpected gamma value. Expected 22 but got %d", gamma)
	}
	if epsilon != 9 {
		t.Errorf("Unexpected expsilon value. Expected 9 but got %d", epsilon)
	}
}

func TestOxygenCo2(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	oxygen := oxygenGenerator(input)
	co2Scrubber := co2ScrubberRating(input)

	if oxygen != 23 {
		t.Errorf("Unexpected Oxygen value. Expected 23 but got %d", oxygen)
	}
	if co2Scrubber != 10 {
		t.Errorf("Unexpected CO2 Scrubber value. Expected 10 but got %d", co2Scrubber)
	}
}
