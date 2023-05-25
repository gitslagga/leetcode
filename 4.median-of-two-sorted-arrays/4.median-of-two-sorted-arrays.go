package solution

import "sort"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var num []int
	num = append(num, nums1...)
	num = append(num, nums2...)
	sort.Ints(num)

	l := len(num)
	if l%2 == 0 {
		return float64(num[l/2-1]+num[l/2]) / 2
	}

	return float64(num[l/2])
}
