package bst

import "fmt"

type BST struct {
	Root *Node
}

func New(val int) *BST {
	return &BST{
		Root: &Node{
			Value: val,
			Left:  nil,
			Right: nil,
		},
	}
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(val int) *Node {
	return &Node{
		Value: val,
		Left:  nil,
		Right: nil,
	}
}

func (b *BST) Insert(val int) error {
	if b == nil {
		return fmt.Errorf("Root node not defined")
	}

	tempRoot := b.Root
	for tempRoot != nil {
		if tempRoot.Value > val {
			// go to the right subchild
			if tempRoot.Right == nil {
				tempRoot.Right = NewNode(val)
				break
			} else {
				tempRoot = tempRoot.Right
			}
		} else {
			// go to left subchild
			if tempRoot.Left == nil {
				tempRoot.Left = NewNode(val)
				break
			} else {
				tempRoot = tempRoot.Left
			}
		}

	}
	fmt.Printf("Added to node: %+v", tempRoot)
	return nil
}
