package simple

import "sort"

// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

// Merge 调用函数进行排序
func Merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// MergeOpt 双指针
func MergeOpt(nums1 []int, m int, nums2 []int, n int) {
	sortNums := make([]int, 0, m+n)
	p, q := 0, 0
	for true {
		if p == m {
			sortNums = append(sortNums, nums2[q:]...)
			break
		}
		if q == n {
			sortNums = append(sortNums, nums1[p:]...)
			break
		}
		if nums1[p] < nums2[q] {
			sortNums = append(sortNums, nums1[p])
			p++
		} else {
			sortNums = append(sortNums, nums2[q])
			q++
		}
	}
	copy(nums1, sortNums)
}
