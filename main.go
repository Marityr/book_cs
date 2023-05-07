package main

import "fmt"

func main() {
	ll := NewList[int]()

	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)
	ll.PushBack(4)
	ll.PushBack(5)
	ll.PushBack(6)

	fmt.Println(ll)
}
