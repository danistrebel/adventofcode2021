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
		{
			"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
			"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}

	for _, test := range tests {
		inputNode := parseTNode(test[0])
		expected := test[1]

		explodedOut := explode(inputNode)

		if inputNode.String() != expected {
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
		intputNode := parseTNode(input)
		expected := test[1]

		splitOut := split(intputNode)

		if intputNode.String() != expected {
			t.Error("Unexpected split output. Expected:", expected, "got", splitOut)
		}
	}
}

func TestAddition(t *testing.T) {
	tests := [][]string{
		{
			"[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
			"[2,9]",
			"[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
		},
		{
			"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
		{
			"[[[[4,3],4],4],[7,[[8,4],9]]]",
			"[1,1]",
			"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
			"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
			"[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]",
		},
	}

	for _, test := range tests {
		inputLeft := test[0]
		inputRight := test[1]
		expected := test[2]

		leftNode := parseTNode(inputLeft)
		rightNode := parseTNode(inputRight)
		addOut := add(leftNode, rightNode).String()

		if addOut != expected {
			t.Error("Unexpected addition output. Expected:\n", expected, "\ngot\n", addOut)
		}
	}
}

func TestMagnitude(t *testing.T) {
	tests := []string{
		"[[1,2],[[3,4],5]]",
		"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
		"[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]",
	}

	results := []int{
		143,
		4140,
		3993,
	}

	for i := 0; i < len(tests); i++ {
		caluclatedMagnitude := calcMagnitude(parseTNode(tests[i]))
		if caluclatedMagnitude != results[i] {
			t.Error("Unexpected Magnutude for", tests[i], "expected", results[i], "got", caluclatedMagnitude)
		}
	}
}

func TestLinesSumSmall(t *testing.T) {
	input := []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
		"[5,5]",
		"[6,6]",
	}

	expected := "[[[[5,0],[7,4]],[5,5]],[6,6]]"

	calculated := addLines(input).String()

	if expected != calculated {
		t.Error("Unexpected sum magnitude. Expected", expected, "got", calculated)
	}
}

func TestLinesSum(t *testing.T) {
	input := []string{
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
		"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
		"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
		"[7,[5,[[3,8],[1,4]]]]",
		"[[2,[2,2]],[8,[8,1]]]",
		"[2,9]",
		"[1,[[[9,3],9],[[9,0],[0,7]]]]",
		"[[[5,[7,4]],7],1]",
		"[[[[4,2],2],6],[8,7]]",
	}

	expected := "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"

	calculated := addLines(input).String()

	if expected != calculated {
		t.Error("Unexpected sum. Expected", expected, "got", calculated)
	}
}

func TestLargestPossibleMagnitude(t *testing.T) {
	input := []string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	}
	expected := 3993
	calculated := largestPossibleMagnitude(input)

	if expected != calculated {
		t.Error("Unexpected max magnitude. Expected", expected, "got", calculated)
	}
}
