package simple

/*
给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的
第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
*/

func StrStr(haystack string,needle string) int {
	n, m := len(haystack), len(needle)
	outer:
		for i := 0; i+m <= n; i++ {
			for  j := range needle {
				if haystack[i+j] != needle[j] {
					continue outer
				}
			}
			return i
		}
	return -1
}
