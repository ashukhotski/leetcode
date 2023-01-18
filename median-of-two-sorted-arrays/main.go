// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).
// See https://leetcode.com/problems/median-of-two-sorted-arrays/
package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n < 1 {
		return 0
	}
	nums := make([]int, n)
	for i := 0; i < len(nums1); i++ {
		nums[i] = nums1[i]
	}
	for j := 0; j < len(nums2); j++ {
		nums[len(nums1)+j] = nums2[j]
	}
	if n == 1 {
		return float64(nums[0])
	}

	for i := 1; i < n; i++ {
		x := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > x {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = x
	}

	if n%2 == 1 {
		return float64(nums[(n-1)/2])
	}

	return float64(nums[n/2]+nums[n/2-1]) / 2
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println("median: ", findMedianSortedArrays(nums1, nums2))
}
