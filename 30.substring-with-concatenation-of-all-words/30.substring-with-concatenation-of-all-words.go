package solution

func FindSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	sLen := len(s)
	wLen := len(words)
	k := len(words[0])
	if sLen < k*wLen {
		return []int{}
	}
	res := []int{}
	mp := make(map[string]int)
	for _, word := range words {
		mp[word]++
	}

	for i := 0; i < k; i++ {
		var count int
		counterMap := make(map[string]int)
		for l, r := i, i; r <= sLen-k; r += k {
			word := s[r : r+k]
			if num, found := mp[word]; found {
				for counterMap[word] >= num {
					counterMap[s[l:l+k]]--
					count--
					l += k
				}
				counterMap[word]++
				count++
			} else {
				for l < r {
					//If the current word is not in the dictionary, move the pointer to the left
					counterMap[s[l:l+k]]--
					count--
					l += k
				}
				l += k
			}
			if count == wLen {
				res = append(res, l)
			}
		}
	}
	return res

}
