package day10

import (
	"testing"
)

func TestSyntaxErrors(t *testing.T) {
	inputs := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}

	corruptionScore := totalCorruptionScore(inputs)
	expectedScore := 26397

	if corruptionScore != expectedScore {
		t.Error("Invalid Corruption. Got", corruptionScore, "expected", expectedScore)
	}
}

func TestAutoClose(t *testing.T) {
	inputs := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}

	autoCloseScore := closingScore(inputs)
	expectedScore := 288957

	if autoCloseScore != expectedScore {
		t.Error("Invalid Corruption. Got", autoCloseScore, "expected", expectedScore)
	}
}
