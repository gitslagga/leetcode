package solution

func IntegerToRoman(num int) string {
	v := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	r := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	re := ""
	for i := 0; i < len(v); i++ {
		for num >= v[i] {
			num -= v[i]
			re += r[i]
		}
	}
	return re
}

func IntToRoman(num int) string {
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hrs := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	ths := []string{"", "M", "MM", "MMM"}

	return ths[num/1000] + hrs[num%1000/100] + tens[num%100/10] + ones[num%10]
}
