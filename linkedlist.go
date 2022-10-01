package main

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T constraints.Ordered] struct {
	head *Node[T]
}

func NewLinkedList[T constraints.Ordered]() LinkedList[T] {
	return LinkedList[T]{
		head: nil,
	}
}

func (a *LinkedList[T]) Append(data T) {
	if a.head == nil {
		a.head = &Node[T]{Data: data, Next: nil}
		return
	}
	cur := a.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &Node[T]{Data: data, Next: nil}
}

func (a *LinkedList[T]) Find(data T) *Node[T] {
	if a.head == nil {
		return nil
	}
	cur := a.head
	for cur.Data == data {
		return cur
	}
	return nil
}
