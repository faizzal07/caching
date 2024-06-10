package cache

import "time"

type ListNode struct {
	prev  *ListNode
	next  *ListNode
	key   string
	value interface{}
	ttl   time.Time
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) PushFront(node *ListNode) {
	node.next = ll.head
	node.prev = nil
	if ll.head != nil {
		ll.head.prev = node
	}
	ll.head = node
	if ll.tail == nil {
		ll.tail = node
	}
}

func (ll *LinkedList) MoveToFront(node *ListNode) {
	if node == ll.head {
		return
	}
	ll.remove(node)
	ll.PushFront(node)
}

func (ll *LinkedList) RemoveLast() *ListNode {
	if ll.tail == nil {
		return nil
	}
	node := ll.tail
	ll.remove(node)
	return node
}

func (ll *LinkedList) remove(node *ListNode) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		ll.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		ll.tail = node.prev
	}
}
