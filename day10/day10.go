package day10

import (
	"adventofcode/utils"
	"fmt"
	"os"
	"sort"
	"strings"
)

type SymbolType uint8

const (
	Unknown SymbolType = iota
	Opening
	Closing
)

func getSymbolType(s string) SymbolType {
	switch s {
	case "{":
		return Opening
	case "[":
		return Opening
	case "<":
		return Opening
	case "(":
		return Opening
	case "}":
		return Closing
	case "]":
		return Closing
	case ">":
		return Closing
	case ")":
		return Closing
	default:
		return Unknown
	}
}

func correspoindingClosingSymbol(openingSymbol string) string {
	switch openingSymbol {
	case "{":
		return "}"
	case "[":
		return "]"
	case "<":
		return ">"
	case "(":
		return ")"
	default:
		return ""
	}
}

func isCorrupted(input string) (bool, string) {

	stack := []string{}

	for _, s := range strings.Split(input, "") {
		switch getSymbolType(s) {
		case Opening:
			stack = append(stack, s)
		case Closing:
			if len(stack) > 0 {
				correctClosing := correspoindingClosingSymbol(stack[len(stack)-1])
				if s == correctClosing {
					stack = stack[:len(stack)-1]
				} else {
					// fmt.Println("expected", correctClosing, "got", s)
					return true, s
				}
			} else {
				return false, ""
			}
		case Unknown:
			fmt.Println("Unknwn Symbol", s)
			os.Exit(-1)
		}
	}
	return false, ""
}

func corruptionScore(input string) int {
	corrupted, firstIllegal := isCorrupted(input)
	if corrupted {
		switch firstIllegal {
		case ")":
			return 3
		case "]":
			return 57
		case "}":
			return 1197
		case ">":
			return 25137
		}
	}
	return 0
}

func totalCorruptionScore(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += corruptionScore(input)
	}
	return sum
}

func autoCloseScore(s string) int {
	switch s {
	case ")":
		return 1
	case "]":
		return 2
	case "}":
		return 3
	case ">":
		return 4
	default:
		fmt.Println("unknown close symbol", s)
		os.Exit(-1)
		return 0
	}
}

func finishClosingSequece(input string) (bool, int) {

	stack := []string{}

	for _, s := range strings.Split(input, "") {
		switch getSymbolType(s) {
		case Opening:
			stack = append(stack, s)
		case Closing:
			if len(stack) > 0 {
				correctClosing := correspoindingClosingSymbol(stack[len(stack)-1])
				if s == correctClosing {
					stack = stack[:len(stack)-1]
				} else {
					// fmt.Println("expected", correctClosing, "got", s)
					return false, -1
				}
			}
		case Unknown:
			fmt.Println("Unknwn Symbol", s)
			os.Exit(-1)
		}
	}

	score := 0
	for i := len(stack) - 1; i >= 0; i-- {
		opening := stack[i]
		closing := correspoindingClosingSymbol(opening)
		score = score*5 + autoCloseScore(closing)

	}
	return true, score
}

func closingScore(inputs []string) int {
	scores := []int{}

	for _, input := range inputs {
		ok, score := finishClosingSequece(input)
		if ok {
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func PrintSolution() {
	lines := utils.ParseLines("./inputs/day10.txt")
	score := totalCorruptionScore(lines)
	fmt.Println("Corrupted Score (Day 1)", score)
	autoCloseScore := closingScore(lines)
	fmt.Println("Auto-close Score (Day 2)", autoCloseScore)
}
