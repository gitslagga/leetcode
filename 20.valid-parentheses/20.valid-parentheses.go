package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsValid(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		if c == '(' {
			stack = append(stack, ')')
		} else if c == '{' {
			stack = append(stack, '}')
		} else if c == '[' {
			stack = append(stack, ']')
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != c {
				return false
			}
		}
	}
	return len(stack) == 0
}
