package bst

type Traversar interface {
	TraversePreOrder(treeRoot *Node) error
	TraverseInOder(treeRoot *Node) error
	TraversePostOrder(treeRoot *Node) error
}

type Recursion struct{}
