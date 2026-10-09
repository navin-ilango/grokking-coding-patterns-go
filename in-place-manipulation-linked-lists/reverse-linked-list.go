package inplacemanipulationlinkedlists

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	next := head.Next

	for next != nil {
		current.Next = prev
		current = next
		prev = current
		next = next.Next
	}
	return current
}
