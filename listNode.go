package main

import "fmt"

type LinkNode struct {
	value any
	next  *LinkNode
}

type LinkedList struct {
	head *LinkNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil}
}

func (ll *LinkedList) InsertNode(value any) {
	var prev *LinkNode
	current := ll.head

	// перебираем список запоминая предыдущий узел пока не дойдем до последнего узла
	for current != nil {
		prev = current
		current = current.next
	}

	// создаем новый узел
	newNode := &LinkNode{value, current}
	// проверяем есть ли предыдущий узел
	if prev == nil {
		// создаем головной узел если если нету предыдущего
		ll.head = newNode
	} else {
		// если есть предшестующий узел то добавляем в него ссылку на новый узел
		prev.next = newNode
	}
}

func (ll *LinkedList) RemoveNode(val int) {
	var prev *LinkNode
	current := ll.head

	// перебираем ноды списка плка не гаькнемся на нужный элемент
	// при этом запоминаем предыдущий элемент и последующий
	for current != nil && current.value != val {
		prev = current
		current = current.next
	}

	// проверяем являетсяли нода головной
	if prev == nil {
		ll.head = current.next
	} else {
		// присваиваем ноде ссылку на следующий элемент за удаленной нодой
		prev.next = current.next
	}
}

func (ll *LinkedList) Reverse() {
	var prev *LinkNode
	current := ll.head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	ll.head = prev
}

func (ll *LinkedList) Find(key int) bool {
	current := ll.head

	for current != nil && key != current.value {
		current = current.next
	}

	return current != nil && key == current.value
}

func (ll *LinkedList) hashLoop() bool {
	if ll.head == nil {
		return false
	}

	slow := ll.head
	fast := ll.head.next

	for fast != nil && fast.next != nil {
		if slow == fast {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}

	return false
}

func (ll *LinkedList) Print() {
	current := ll.head

	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
