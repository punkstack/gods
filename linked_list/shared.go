package linked_list

type Node struct {
	value interface{}
	next  *Node
}

func NewNode(v interface{}) *Node {
	return &Node{value: v, next: nil}
}
