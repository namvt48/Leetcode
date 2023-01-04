// func isAnagram(s string, t string) bool {
// 	m := make(map[string]int)

// 	for _, v := range s {
// 		m[string(v)]++
// 	}

// 	for _, v := range t {
// 		m[string(v)]--
// 		if m[string(v)] == 0 {
// 			delete(m, string(v))
// 		}
// 	}
// 	if len(m) == 0 {
// 		return true
// 	}
// 	return false
// }

func isAnagram(s string, t string) bool {
	m := [128]int{}

	for _, v := range s {
		a := v - '0'
		m[int(a)]++
	}

	for _, v := range t {
		a := v - '0'
		m[int(a)]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

