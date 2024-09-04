package mergetwosortedlist

import (
	"reflect"
	"testing"
)

// Helper function to create a linked list from a slice
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head
	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

// Helper function to convert a linked list to a slice
func listToSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	return vals
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{list1: []int{}, list2: []int{}, want: []int{}},                               // both empty lists
		{list1: []int{1, 2, 4}, list2: []int{}, want: []int{1, 2, 4}},                 // one list empty
		{list1: []int{}, list2: []int{1, 3, 4}, want: []int{1, 3, 4}},                 // one list empty
		{list1: []int{1, 2, 4}, list2: []int{1, 3, 4}, want: []int{1, 1, 2, 3, 4, 4}}, // normal case
		{list1: []int{5}, list2: []int{6}, want: []int{5, 6}},                         // both single element
	}

	for _, tt := range tests {
		list1 := createList(tt.list1)
		list2 := createList(tt.list2)
		got := mergeTwoLists(list1, list2)

		if !reflect.DeepEqual(listToSlice(got), tt.want) {
			t.Errorf("mergeTwoLists(%v, %v) = %v, want %v", tt.list1, tt.list2, listToSlice(got), tt.want)
		}
	}
}
