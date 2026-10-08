package fastslowpointers

type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedList represents a linked list
type LinkedList struct {
	Head *ListNode
}

func detectCycle(head *ListNode) bool {

	slowPtr := head
	fastPtr := head.Next

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if slowPtr == fastPtr {
			return true
		}
	}

	return false
}
