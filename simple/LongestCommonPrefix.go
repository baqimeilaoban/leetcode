package simple

// 最长公共前缀

// LongestCommonPrefix 横向扫描法
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			return prefix
		}
	}
	return prefix
}

// 获取str1和str2中公共前缀
func lcp(str1, str2 string) string {
	minLen := min(str1, str2)
	index := 0
	for index < minLen && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

// 求两字符串中长度较小的长度
func min(str1, str2 string) int {
	if len(str1) < len(str2) {
		return len(str1)
	}
	return len(str2)
}

// LongestCommonPrefixOpt 纵向扫描
func LongestCommonPrefixOpt(strs []string) string {
	var res string
	if len(strs) == 0 {
		return res
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				res = strs[0][:i]
				return res
			}
		}
	}
	// 防止传过来的str数组中只有一个字符串
	return strs[0]
}
