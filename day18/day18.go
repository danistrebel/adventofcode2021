package day18

import (
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

func parseTNode(input string) TNode {
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
			fmt.Println(value)
			currentNode.value = value
		}

	}

	return *currentNode
}

func findNextLeft(t *TNode) *TNode {
	currentNode := t
	for currentNode.parent != nil {
		if currentNode.parent.right == currentNode {
			return currentNode.parent.left
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

func isNumberPair(t TNode) bool {
	return t.left != nil && t.left.value >= 0 && t.right != nil && t.right.value >= 0
}

func explode(t TNode) TNode {
	root := &t
	currentNode := root
	currentDepth := 0

	for {
		if currentDepth >= 4 && isNumberPair(*currentNode) {
			//explode
			leftValue := currentNode.left.value
			rightValue := currentNode.right.value

			currentNode.left = nil
			currentNode.right = nil
			currentNode.value = 0

			// add left number to next left
			if nextLeft := findNextLeft(currentNode); nextLeft != nil {
				nextLeft.value += leftValue
			}

			// add right number to next left
			if nextRight := findNextRight(currentNode); nextRight != nil {
				nextRight.value += rightValue
			}
			break
		} else if currentNode.left != nil {
			currentDepth += 1
			currentNode = currentNode.left
		} else if currentNode.right != nil {
			currentDepth += 1
			currentNode = currentNode.right
		} else {
			for currentNode.parent != nil {
				if currentNode.parent.left == currentNode {
					currentNode = currentNode.parent.right
					break
				}
				currentDepth -= 1
				currentNode = currentNode.parent
			}

			if currentNode.parent == nil {
				break //back at root
			}
		}
	}
	return *root
}

func split(t TNode) TNode {
	root := &t
	currentNode := root
	fmt.Println(currentNode)

	for {
		if currentNode.value >= 10 {
			//explode
			leftValue := int(math.Floor(float64(currentNode.value) / 2.0))
			rightValue := int(math.Ceil(float64(currentNode.value) / 2.0))

			currentNode.left = &TNode{currentNode, nil, nil, leftValue}
			currentNode.right = &TNode{currentNode, nil, nil, rightValue}
			currentNode.value = -1

			fmt.Println(currentNode)

			break
		} else if currentNode.left != nil {
			currentNode = currentNode.left
		} else if currentNode.right != nil {
			currentNode = currentNode.right
		} else {
			for currentNode.parent != nil {
				if currentNode.parent.left == currentNode {
					currentNode = currentNode.parent.right
					break
				}
				currentNode = currentNode.parent
			}

			if currentNode.parent == nil {
				break //back at root
			}
		}
	}
	return *root
}

func (tn TNode) String() string {
	if tn.value != -1 {
		return fmt.Sprint(tn.value)
	}
	return fmt.Sprintf("[%s,%s]", tn.left, tn.right)
}
