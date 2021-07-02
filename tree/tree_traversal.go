package tree

type Traversar interface {
	TraversePreOrder(treeRoot *Node) error
	TraverseInOder(treeRoot *Node) error
	TraversePostOrder(treeRoot *Node) error
}

type Recursion struct{}

func (r *Recursion) TraversePreOrder(treeRoot *tree.Node) error {
}

func (r *Recursion) TraverseInOder(treeRoot *tree.Node) error {
	panic("not implemented") // TODO: Implement
}

func (r *Recursion) TraversePostOrder(treeRoot *tree.Node) error {
	panic("not implemented") // TODO: Implement
}

type Loop struct{}

func (l *Loop) TraversePreOrder(treeRoot *tree.Node) error {
	panic("not implemented") // TODO: Implement
}

func (l *Loop) TraverseInOder(treeRoot *tree.Node) error {
	panic("not implemented") // TODO: Implement
}

func (l *Loop) TraversePostOrder(treeRoot *tree.Node) error {
	panic("not implemented") // TODO: Implement
}
