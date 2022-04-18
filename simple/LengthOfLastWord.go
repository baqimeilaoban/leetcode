package simple

import "strings"

// 给你一个字符串 s，由若干 单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

func LengthOfLastWord(s string) int {
	strs := strings.Split(s, " ")
	for i := len(strs) - 1; i >= 0; i-- {
		if len(strs[i]) != 0 {
			return len(strs[i])
		}
	}
	return 0
}

func LengthOfLastWordOpt(s string) int {
	var res int
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		index--
		res++
	}
	return res
}
