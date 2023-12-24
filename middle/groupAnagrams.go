package middle

import "sort"

// 字母易位词分组
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
/*
示例1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

示例2:
输入: strs = [""]
输出: [[""]]

示例3:
输入: strs = ["a"]
输出: [["a"]]
*/

func groupAnagramsMethod1(strs []string) [][]string {
	// 构建哈希表存储相同的词
	mp := make(map[string][]string, 0)
	for _, str := range strs {
		bs := []byte(str)
		// 单个str进行排序，将诸如ant,tan,atn均排序为ant
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] > bs[j]
		})
		sortStr := string(bs)
		mp[sortStr] = append(mp[sortStr], str)
	}
	var res [][]string
	for _, sortStrs := range mp {
		res = append(res, sortStrs)
	}
	return res
}

func groupAnagramsMethod2(strs []string) [][]string {
	// 计数法 每个str均有小写字母组成，极端情况下map的key是一个长度为26的数组
	groupStrMp := make(map[[26]int][]string)
	for _, str := range strs {
		ctnBs := [26]int{}
		for _, tmpByte := range str {
			// ant与tan经过计算后，在map中的位置一致，所以直接相加即可
			ctnBs[tmpByte-'a']++
		}
		groupStrMp[ctnBs] = append(groupStrMp[ctnBs], str)
	}
	var res [][]string
	for _, strings := range groupStrMp {
		res = append(res, strings)
	}
	return res
}
