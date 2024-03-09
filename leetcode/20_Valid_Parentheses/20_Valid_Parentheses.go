package leetcode

type Stack struct {
	s []int
}

func (s *Stack) Push(value int) {
	s.s = append(s.s, value)
}

func (s *Stack) Pop() int {
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) Length() int {
	length := len(s.s)
	return length
}

// time complexity is O(n), space complexity is O(n)
func isValid(s string) bool {
	stack := new(Stack)
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack.Push(int(char))
		} else if stack.Length() == 0 {
			return false
		} else {
			openParenthese := stack.Pop()
			expectedParenthese := '('
			switch char {
				case ')':
					expectedParenthese = '('
				case '}':
					expectedParenthese = '{'
				case ']':
					expectedParenthese = '['
			}

			if rune(openParenthese) != expectedParenthese {
				return false
			}
		}
	}

	if stack.Length() == 0 {
		return true
	}

	return false
}
