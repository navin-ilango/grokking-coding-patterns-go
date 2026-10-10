package stacks

func removeDuplicates(s string) string {

	stack := make([]rune, 0, len(s))

	for _, c := range s {
		if len(stack) > 0 && stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return string(stack)
}
