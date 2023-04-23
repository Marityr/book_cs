package main

func main() {
	ll := NewLinkedList()

	ll.InsertNode(1)
	ll.InsertNode(2)
	ll.InsertNode(3)
	ll.InsertNode(4)
	ll.InsertNode(5)
	ll.InsertNode(6)

	ll.Reverse()
	ll.Print()
}
