package solution

func RomanToInteger(s string) int {
	v := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	r := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	re := 0

	for i := 0; i < len(r); i++ {
		prefix := r[i]
		for len(s) >= len(prefix) && s[0:len(prefix)] == prefix {
			s = s[len(prefix):]
			re += v[i]
		}
	}
	return re
}
