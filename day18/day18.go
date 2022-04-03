package day18

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strconv"
)

type TNode struct {
	parent *TNode
	left   *TNode
	right  *TNode
	value  int
}

func (tn TNode) String() string {
	if tn.isLeaf() {
		return fmt.Sprint(tn.value)
	}
	return fmt.Sprintf("[%s,%s]", tn.left, tn.right)
}

func (tn TNode) isLeaf() bool {
	return tn.left == nil && tn.right == nil
}

func parseTNode(input string) *TNode {
	var currentNode *TNode
	currentNode = &TNode{nil, nil, nil, -1}

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '[':
			{
				nextNode := TNode{currentNode, nil, nil, -1}
				currentNode.left = &nextNode
				currentNode = &nextNode
			}
		case ',':
			{
				nextNode := TNode{currentNode.parent, nil, nil, -1}
				currentNode.parent.right = &nextNode
				currentNode = &nextNode
			}
		case ']':
			currentNode = currentNode.parent
		default:
			numberLength := 1
			value := -1
			for {
				parsed, err := strconv.ParseInt(input[i:i+numberLength], 10, 64)
				if err != nil {
					break
				} else {
					value = int(parsed)
					numberLength++
				}
			}
			i += (numberLength - 2)
			currentNode.value = value
		}

	}

	return currentNode
}

func findNextLeft(t *TNode) *TNode {
	currentNode := t
	for currentNode.parent != nil {
		if currentNode.parent.right == currentNode {
			currentNode = currentNode.parent.left
			for currentNode.right != nil {
				currentNode = currentNode.right
			}
			return currentNode
		}
		currentNode = currentNode.parent
	}
	return nil
}

func findNextRight(t *TNode) *TNode {
	currentNode := t
	for currentNode.parent != nil {
		if currentNode.parent.left == currentNode {
			currentNode = currentNode.parent.right
			for currentNode != nil && currentNode.left != nil {
				currentNode = currentNode.left
			}
			return currentNode
		}
		currentNode = currentNode.parent
	}
	return nil
}

func findNodeAtDepth(t *TNode, depthLeft int) *TNode {
	if depthLeft == 0 && !t.isLeaf() {
		return t
	}

	if depthLeft <= 0 {
		return nil
	}

	if t.left != nil {
		if foundLeft := findNodeAtDepth(t.left, depthLeft-1); foundLeft != nil {
			return foundLeft
		}
	}

	if t.right != nil {
		if foundRight := findNodeAtDepth(t.right, depthLeft-1); foundRight != nil {
			return foundRight
		}
	}
	return nil
}

func findValueGreaterEqualThan(t *TNode, v int) *TNode {
	if t.isLeaf() && t.value >= v {
		return t
	}
	if t.left != nil {
		if leftFound := findValueGreaterEqualThan(t.left, v); leftFound != nil {
			return leftFound
		}
	}
	if t.right != nil {
		if rightFound := findValueGreaterEqualThan(t.right, v); rightFound != nil {
			return rightFound
		}
	}
	return nil
}

func explode(t *TNode) bool {
	if nodeToExplode := findNodeAtDepth(t, 4); nodeToExplode != nil {
		leftValue := nodeToExplode.left.value
		rightValue := nodeToExplode.right.value

		nodeToExplode.left = nil
		nodeToExplode.right = nil
		nodeToExplode.value = 0

		// add left number to next left
		if nextLeft := findNextLeft(nodeToExplode); nextLeft != nil {
			nextLeft.value += leftValue
		}

		// add right number to next left
		if nextRight := findNextRight(nodeToExplode); nextRight != nil {
			nextRight.value += rightValue
		}
		return true
	}
	return false
}

func split(t *TNode) bool {
	if valueToSplit := findValueGreaterEqualThan(t, 10); valueToSplit != nil {
		leftValue := int(math.Floor(float64(valueToSplit.value) / 2.0))
		rightValue := int(math.Ceil(float64(valueToSplit.value) / 2.0))

		valueToSplit.left = &TNode{valueToSplit, nil, nil, leftValue}
		valueToSplit.right = &TNode{valueToSplit, nil, nil, rightValue}
		valueToSplit.value = -1
		return true
	}
	return false
}

func add(l *TNode, r *TNode) *TNode {
	additionRoot := TNode{nil, l, r, -1}
	l.parent = &additionRoot
	r.parent = &additionRoot

	splited := true
	exploded := true

	for splited || exploded {
		for exploded {
			exploded = explode(&additionRoot)
		}
		splited = split(&additionRoot)
		exploded = explode(&additionRoot)
	}

	return &additionRoot
}

func calcMagnitude(t *TNode) int {
	if t.value != -1 {
		return t.value
	} else {
		return 3*calcMagnitude(t.left) + 2*calcMagnitude(t.right)
	}
}

func addLines(lines []string) *TNode {
	prevSum := parseTNode(lines[0])

	for i := 1; i < len(lines); i++ {
		nextLine := parseTNode(lines[i])
		prevSum = add(prevSum, nextLine)
	}

	return prevSum
}

func largestPossibleMagnitude(lines []string) int {
	largest := 0
	for i := 0; i < len(lines)-1; i++ {
		for j := 0; j < len(lines); j++ {

			if j != i {
				iNode := parseTNode(lines[i])
				jNode := parseTNode(lines[j])

				if mag := calcMagnitude(add(iNode, jNode)); mag > largest {
					largest = mag
				}
			}
		}
	}
	return largest
}

func PrintSolution() {
	lines := utils.ParseLines("./inputs/day18.txt")
	magnitude := calcMagnitude(addLines(lines))
	fmt.Println("Magnitude (Day 1)", magnitude)
	MaxMagnitude := largestPossibleMagnitude(lines)
	fmt.Println("MaxMagnitude (Day 2)", MaxMagnitude)
}
