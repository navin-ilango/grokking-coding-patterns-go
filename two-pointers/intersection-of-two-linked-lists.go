package main

import "fmt"

// Definition for a Linked List node

type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedList represents a linked list
type LinkedList struct {
	Head *ListNode
}

// IDEA: Switch Paths
// Instead of counting lengths or using a hash set,
// you can make two pointers travel the exact same
// total distance by switching to the other list
// when they reach the end

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	ptrA, ptrB := headA, headB
	for ptrA != ptrB {
		if ptrA == nil {
			ptrA = headB
		} else {
			ptrA = ptrA.Next
		}
		if ptrB == nil {
			ptrB = headA
		} else {
			ptrB = ptrB.Next
		}
	}
	return ptrA

}

// NewLinkedList creates a linked list from a slice of integers
func NewLinkedList(values []int) *LinkedList {
	ll := &LinkedList{}
	if len(values) > 0 {
		ll.createLinkedList(values)
	}
	return ll
}

// createLinkedList initializes the linked list from a slice of values
func (ll *LinkedList) createLinkedList(values []int) {
	if len(values) == 0 {
		ll.Head = nil
		return
	}

	ll.Head = &ListNode{Val: values[0]}
	current := ll.Head
	for _, value := range values[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}
}

// display prints the linked list
func display(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("None")
}
