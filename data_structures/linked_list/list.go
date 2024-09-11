package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertAtFront(head *ListNode, data int) *ListNode {
	newHead := &ListNode{Val: data}
	if head == nil {
		return newHead
	}
	newHead.Next = head
	return newHead
}

func InsertAtEnd(head *ListNode, data int) *ListNode {
	newNode := &ListNode{Val: data}
	if head == nil {
		return newNode
	}

	last := head
	for last.Next != nil {
		last = last.Next
	}

	last.Next = newNode
	return head
}

func Traverse(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
