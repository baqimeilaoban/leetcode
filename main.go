package main

import (
	"fmt"

	"github.com/baqimeilaoban/leetcode/simple"
)

func main() {
	nums1 := []int{1}
	nums2 := []int{}
	m := 1
	n := 0
	simple.MergeOpt(nums1, m, nums2, n)
	fmt.Println(nums1)
}
