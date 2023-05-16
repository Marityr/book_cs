package main

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewList[T any]() *List[T] {
	return &List[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (ll *List[T]) PushBack(value T) {
	if ll.head == nil {
		ll.head = &Node[T]{
			value: value,
			next:  nil,
		}
		ll.tail = ll.head
		return
	}

	if ll.tail != nil {
		ll.tail.next = &Node[T]{
			value: value,
			next:  nil,
		}
		ll.tail = ll.tail.next
		return
	}

	ll.size++
}

func (ll *List[T]) PopFront() {
	if ll.head == nil {
		return
	} else if ll.head.next != nil {
		ll.head = ll.head.next
		ll.size--
	}
}

// func (ll *LinkedList) RemoveNode(val int) {
// 	var prev *LinkNode
// 	current := ll.head

// 	// перебираем ноды списка плка не гаькнемся на нужный элемент
// 	// при этом запоминаем предыдущий элемент и последующий
// 	for current != nil && current.value != val {
// 		prev = current
// 		current = current.next
// 	}

// 	// проверяем являетсяли нода головной
// 	if prev == nil {
// 		ll.head = current.next
// 	} else {
// 		// присваиваем ноде ссылку на следующий элемент за удаленной нодой
// 		prev.next = current.next
// 	}
// }

// func (ll *LinkedList) Reverse() {
// 	var prev *LinkNode
// 	current := ll.head

// 	for current != nil {
// 		next := current.next
// 		current.next = prev
// 		prev = current
// 		current = next
// 	}

// 	ll.head = prev
// }

// func (ll *LinkedList) Find(key int) bool {
// 	current := ll.head

// 	for current != nil && key != current.value {
// 		current = current.next
// 	}

// 	return current != nil && key == current.value
// }

// func (ll *LinkedList) hashLoop() bool {
// 	if ll.head == nil {
// 		return false
// 	}

// 	slow := ll.head
// 	fast := ll.head.next

// 	for fast != nil && fast.next != nil {
// 		if slow == fast {
// 			return true
// 		}
// 		slow = slow.next
// 		fast = fast.next.next
// 	}

// 	return false
// }

func (ll *List[T]) String() string {
	var bulder strings.Builder
	current := ll.head

	for current != nil {
		if current != nil && current != ll.head {
			bulder.WriteString(" -> ")
		}
		str := fmt.Sprintf("%v", current.value)
		bulder.WriteString(str)
		current = current.next
	}

	return bulder.String()
}
