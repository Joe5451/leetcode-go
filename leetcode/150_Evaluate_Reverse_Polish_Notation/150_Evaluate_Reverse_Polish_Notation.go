package leetcode

import "strconv"

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
func evalRPN(tokens []string) int {
	stack := new(Stack)

	for _, char := range tokens {
		num, err := strconv.Atoi(char)

		if err == nil {
			stack.Push(num)
		} else {
			num1 := stack.Pop()
			num2 := stack.Pop()
			switch char {
				case "+":
					stack.Push(num2 + num1)
				case "-":
					stack.Push(num2 - num1)
				case "*":
					stack.Push(num2 * num1)
				case "/":
					stack.Push(num2 / num1)
			}
		}
	}

	return stack.Pop()
}
