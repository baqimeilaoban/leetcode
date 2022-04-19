package simple

import "strconv"

// 给你两个二进制字符串，返回它们的和（用二进制表示）。

func AddBinary(a string, b string) string {
	res := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		res = strconv.Itoa(carry%2) + res
		carry /= 2
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
