package solution

func LongestPalindrome(s string) string {
	l := len(s)
	ret := string(s[0])
	for i := 0; i < l; i++ {
		// * * * xx * * *
		if i < l-1 && s[i] == s[i+1] {
			j := 1
			for i-j >= 0 && i+j+1 < l && s[i-j] == s[i+j+1] {
				j++
			}
			if len(ret) < j*2 {
				ret = s[i-j+1 : i+j+1]
			}
		}
		// * * * xyx * * *
		if i > 0 && i < l-1 && s[i-1] == s[i+1] {
			j := 1
			for i-j-1 >= 0 && i+j+1 < l && s[i-j-1] == s[i+j+1] {
				j++
			}
			if len(ret) < j*2+1 {
				ret = s[i-j : i+j+1]
			}
		}
	}
	return ret
}
