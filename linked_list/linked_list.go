package linked_list

import "fmt"

type LinkedList struct {
	Head   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil, length: 0}
}

func (ll *LinkedList) InsertAtStart(v interface{}) {
	temp := ll.Head
	ll.Head = NewNode(v)
	ll.Head.next = temp
	ll.length++
}

func (ll *LinkedList) PritnLinkedList() {
	node := ll.Head
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func (ll *LinkedList) InsertAtEnd(v interface{}) {
	currentNode := ll.Head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = NewNode(v)
	ll.length++
}

func (ll *LinkedList) ReverseLinkedList() {
	if ll.Head == nil {
		return
	}
	currentNode := ll.Head
	var prevNode *Node
	for currentNode != nil {
		temp := currentNode
		currentNode = currentNode.next
		temp.next = prevNode
		prevNode = temp
	}
	ll.Head = prevNode
}

func (ll *LinkedList) ReverseKNodes(n int) {

}

func (ll *LinkedList) FindDuplicates() []*Node {
	return []*Node{}
}
