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

func TestExplode(t *testing.T) {
	tests := [][]string{
		{
			"[[[[[9,8],1],2],3],4]",
			"[[[[0,9],2],3],4]",
		},
		{
			"[7,[6,[5,[4,[3,2]]]]]",
			"[7,[6,[5,[7,0]]]]",
		},
		{
			"[[6,[5,[4,[3,2]]]],1]",
			"[[6,[5,[7,0]]],3]",
		},
		{
			"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}

	for _, test := range tests {
		input := test[0]
		expected := test[1]

		explodedOut := explode(parseTNode(input)).String()

		if explodedOut != expected {
			t.Error("Unexpected explode output. Expected:", expected, "got", explodedOut)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := [][]string{
		{
			"[[[[0,7],4],[15,[0,13]]],[1,1]]",
			"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		},
		{
			"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
			"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
		},
	}

	for _, test := range tests {
		input := test[0]
		expected := test[1]

		splitOut := split(parseTNode(input)).String()

		if splitOut != expected {
			t.Error("Unexpected split output. Expected:", expected, "got", splitOut)
		}
	}
}
