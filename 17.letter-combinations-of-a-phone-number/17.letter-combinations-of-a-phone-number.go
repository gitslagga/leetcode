package solution

func LetterCombinations(digits string) []string {
	dic := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var res []string
	var fn func(int, []rune)
	fn = func(n int, arr []rune) {
		if len(arr) == len(digits) {
			if len(arr) > 0 {
				str := string(arr)
				res = append(res, str)
			}
			return
		}

		for i := n; i < len(digits); i++ {
			index := int(rune(digits[i]) - '0')
			runes := []rune(dic[index])
			for _, v := range runes {
				arr = append(arr, v)
				fn(i+1, arr)
				arr = arr[:len(arr)-1]
			}
		}
	}
	fn(0, []rune{})

	return res
}
