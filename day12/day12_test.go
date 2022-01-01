package day12

import (
	"testing"
)

func TestCaveSize(t *testing.T) {
	big := Cave{"BIGCAVE", []string{}}
	small := Cave{"kj", []string{"BIGCAVE"}}
	if !isBigCave(big.id) {
		t.Error("Expected this to be a big cave")
	}
	if isBigCave(small.id) {
		t.Error("Expected this to be a small cave")
	}
}

func TestCaveLayoutParser(t *testing.T) {
	links := []string{
		"dc-end",
		"HN-start",
		"start-kj",
		"dc-start",
		"dc-HN",
		"LN-dc",
		"HN-end",
		"kj-sa",
		"kj-HN",
		"kj-dc",
	}

	caves := parseCaveLayout(links)

	if len(caves) != 7 {
		t.Error("Unexpected size of cave map", len(caves))
	}
}

func TestPaths(t *testing.T) {
	links := []string{
		"dc-end",
		"HN-start",
		"start-kj",
		"dc-start",
		"dc-HN",
		"LN-dc",
		"HN-end",
		"kj-sa",
		"kj-HN",
		"kj-dc",
	}

	paths := findNumberOfPaths(links)
	expectedNumberOfPaths := 19

	if paths != expectedNumberOfPaths {
		t.Error("Unexpected number of paths. Got", paths, "expected", expectedNumberOfPaths)
	}
}
