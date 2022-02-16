package main

import (
	"fmt"
	"gods/linked_list"
)

func main() {
	fmt.Println("Main")
	ll := linked_list.NewLinkedList()
	ll.InsertAtStart(1)
	ll.InsertAtStart(2)
	ll.InsertAtStart(3)
	ll.InsertAtStart(4)
	ll.PritnLinkedList()
	ll.ReverseLinkedList()
	ll.PritnLinkedList()
}
