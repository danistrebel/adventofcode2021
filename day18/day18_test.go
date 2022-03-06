package day18

import (
	"testing"
)

func TestNodeParser(t *testing.T) {

	inputs := []string{
		"[[[[[9,8],1],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]",
		"[[6,[5,[4,[3,2]]]],1]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
	}

	for _, input := range inputs {
		parserOut := parseTNode(input)
		if input != parserOut.String() {
			t.Error("Unexpected parser output. Expected:", input, "got", parserOut.String())
		}
	}
}
