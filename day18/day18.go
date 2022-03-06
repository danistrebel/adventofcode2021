package day18

import "fmt"

type TNode struct {
	parent *TNode
	left   *TNode
	right  *TNode
	value  int
}

func parseTNode(input string) *TNode {
	var currentNode *TNode
	currentNode = &TNode{nil, nil, nil, -1}

	for i := 0; i < len(input); i++ {
		fmt.Print(string(input[i]))
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
			currentNode.value = int(input[i] - '0')
		}

	}

	return currentNode
}

func (tn TNode) String() string {
	if tn.value != -1 {
		return fmt.Sprint(tn.value)
	}
	return fmt.Sprintf("[%s,%s]", tn.left, tn.right)
}
