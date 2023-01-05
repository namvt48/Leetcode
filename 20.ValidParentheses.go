func isValid(s string) bool {
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]byte, 0)

	for _, v := range []byte(s) { // "[]{}"

		fmt.Println(stack)
		p, oke := pairs[v]

		if !oke {
			stack = append(stack, v)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if p != stack[len(stack)-1] {
			return false
		}

		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}