package solution

import "math"

func Divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return dividend
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		} else {
			return 0
		}
	}
	sign := true
	if dividend > 0 {
		sign = !sign
		dividend = -dividend
	}
	if divisor > 0 {
		sign = !sign
		divisor = -divisor
	}
	result := 0
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)>>1 // watch out for overflow, and don't use division
		if quickAdd(divisor, mid, dividend) {
			result = mid
			if mid == math.MaxInt32 { // watch out for overflow
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if !sign {
		result = -result
	}
	return result
}

// compute a * m < r true or false
func quickAdd(a int, m int, r int) bool {
	result, add := 0, a
	for m > 0 {
		if m&1 > 0 {
			if result < r-add {
				return false
			}
			result += add
		}
		if m != 1 {
			if add < r-add {
				return false
			}
			add += add
		}
		m >>= 1
	}
	return true
}
