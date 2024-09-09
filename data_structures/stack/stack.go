package stack

type Stack struct {
	elements []int
}

func (s *Stack) Push(item int) {
	s.elements = append(s.elements, item)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}

	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:(len(s.elements) - 1)]
	return top, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Top() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}

	return s.elements[len(s.elements)-1], true
}
