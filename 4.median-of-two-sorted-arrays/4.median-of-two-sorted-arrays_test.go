package solution

import (
	"testing"
)

func TestFindMedianSortedArrays1(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	expectResult := 2.00000

	result := FindMedianSortedArrays(nums1, nums2)
	if result != expectResult {
		t.Errorf("return wrong result, expect:%v, got:%v", expectResult, result)
	}
}

func TestFindMedianSortedArrays2(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	expectResult := 2.50000

	result := FindMedianSortedArrays(nums1, nums2)
	if result != expectResult {
		t.Errorf("return wrong result, expect:%v, got:%v", expectResult, result)
	}
}
