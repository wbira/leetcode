package parentheses

type Stack struct {
	elements []rune
}

func (s *Stack) Push(item rune) {
	s.elements = append(s.elements, item)
}

func (s *Stack) Pop() (rune, bool) {
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

func isPair(e1, e2 rune) bool {
	if e1 == '(' && e2 == ')' || e2 == '(' && e1 == ')' {
		return true
	}

	if e1 == '[' && e2 == ']' || e2 == '[' && e1 == ']' {
		return true
	}

	if e1 == '{' && e2 == '}' || e2 == '{' && e1 == '}' {
		return true
	}

	return false
}

func IsValid(s string) bool {
	stack := Stack{}
	if len(s)%2 == 1 {
		return false
	}
	for _, v := range s {
		if v == ')' || v == ']' || v == '}' {
			el, _ := stack.Pop()
			if !isPair(el, v) {
				return false
			}
		} else {
			stack.Push(v)
		}
	}

	return stack.IsEmpty()
}
