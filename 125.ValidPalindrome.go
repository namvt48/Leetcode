func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1

	for (s[left] < 97 || s[left] > 122) && (s[left] < 48 || s[left] > 57) {
		left++

		if left == len(s) {
			return false
		}
	}
	for (s[right] < 97 || s[right] > 122) && (s[right] < 48 || s[right] > 57) {
		right--

		if right < 0 {
			return false
		}
	}

	for left < right {
		for (s[left] < 97 || s[left] > 122) && (s[left] < 48 || s[left] > 57) {
			left++
		}
		for (s[right] < 97 || s[right] > 122) && (s[right] < 48 || s[right] > 57) {
			right--
		}

		// fmt.Println(left, right)

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
