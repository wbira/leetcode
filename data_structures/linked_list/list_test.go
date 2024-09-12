package linkedlist

import (
	"reflect"
	"testing"
)

func TestInsertAtFront(t *testing.T) {
	var head *ListNode

	// Insert at front when the list is empty
	head = InsertAtFront(head, 10)
	if head == nil || head.Val != 10 {
		t.Errorf("Expected 10 at the head, but got %v", head)
	}

	// Insert another element at front
	head = InsertAtFront(head, 20)
	if head == nil || head.Val != 20 {
		t.Errorf("Expected 20 at the head, but got %v", head)
	}

	// Verify the second node
	if head.Next == nil || head.Next.Val != 10 {
		t.Errorf("Expected 10 as the second node, but got %v", head.Next)
	}
}

func TestInsertAtEnd(t *testing.T) {
	var head *ListNode

	// Insert at end when the list is empty
	head = InsertAtEnd(head, 10)
	if head == nil || head.Val != 10 {
		t.Errorf("Expected 10 at the head, but got %v", head)
	}

	// Insert another element at end
	head = InsertAtEnd(head, 20)
	if head.Next == nil || head.Next.Val != 20 {
		t.Errorf("Expected 20 at the end, but got %v", head.Next)
	}
}

func TestTraverse(t *testing.T) {
	var head *ListNode

	// Insert multiple nodes to create the list
	head = InsertAtEnd(head, 10)
	head = InsertAtEnd(head, 20)
	head = InsertAtEnd(head, 30)

	// Traverse should return the values in the correct order
	result := Traverse(head)
	expected := []int{10, 20, 30}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestDeleteNode(t *testing.T) {
	var head *ListNode

	// Insert multiple nodes to create the list
	head = InsertAtEnd(head, 10)
	head = InsertAtEnd(head, 20)
	head = InsertAtEnd(head, 30)

	// Delete a node in the middle
	head = DeleteNode(head, 20)
	result := Traverse(head)
	expected := []int{10, 30}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Delete the head node
	head = DeleteNode(head, 10)
	result = Traverse(head)
	expected = []int{30}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Delete the last remaining node
	head = DeleteNode(head, 30)
	if head != nil {
		t.Errorf("Expected empty list, but got %v", Traverse(head))
	}
}
