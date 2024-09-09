package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := Stack{}
	s.Push(10)

	if s.Size() != 1 {
		t.Errorf("Expected size 1, but got %d", s.Size())
	}

	top, _ := s.Top()
	if top != 10 {
		t.Errorf("Expected top element 10, but got %d", top)
	}
}

func TestPop(t *testing.T) {
	s := Stack{}
	s.Push(20)

	value, success := s.Pop()
	if !success {
		t.Errorf("Expected success, but got failure")
	}

	if value != 20 {
		t.Errorf("Expected popped value 20, but got %d", value)
	}

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty, but it's not")
	}
}

func TestIsEmpty(t *testing.T) {
	s := Stack{}

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty initially")
	}

	s.Push(30)
	if s.IsEmpty() {
		t.Errorf("Expected stack to not be empty after push")
	}
}

func TestSize(t *testing.T) {
	s := Stack{}

	if s.Size() != 0 {
		t.Errorf("Expected initial size 0, but got %d", s.Size())
	}

	s.Push(40)
	s.Push(50)

	if s.Size() != 2 {
		t.Errorf("Expected size 2, but got %d", s.Size())
	}
}

func TestTop(t *testing.T) {
	s := Stack{}
	_, success := s.Top()

	if success {
		t.Errorf("Expected no top element in an empty stack")
	}

	s.Push(60)
	value, success := s.Top()
	if !success || value != 60 {
		t.Errorf("Expected top value 60, but got %d", value)
	}
}
