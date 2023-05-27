package solution

import (
	"math"
)

func MyAtoi(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s)
	// Remove any additional spaces before and after the given string
	for i < n && s[i] == ' ' {
		i++
	}
	// Handling numbers with +/- sign
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0')  // ASCII of '0' = 48
		if sign*abs < math.MinInt32 { // Check if result exceeds MinInt32 or MaxInt32
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs
}
