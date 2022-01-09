package day14

import "testing"

func TestPolyRules(t *testing.T) {
	inputs := []string{
		"NNCB",
		"",
		"CH -> B",
		"HH -> N",
		"CB -> H",
		"NH -> C",
		"HB -> C",
		"HC -> B",
		"HN -> C",
		"NN -> C",
		"BH -> H",
		"NC -> B",
		"NB -> B",
		"BN -> B",
		"BB -> N",
		"BC -> B",
		"CC -> N",
		"CN -> C",
	}

	diff := polymerSimulation(inputs, 10)

	expectedDiff := 1588

	if diff != expectedDiff {
		t.Error("Unexpected Count diff. Got", diff, "expected", expectedDiff)
	}
}
