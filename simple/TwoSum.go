package simple

// 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标
/*
示例1:
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例2:
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例3:
输入：nums = [3,3], target = 6
输出：[0,1]
*/

// 方法一: 双重for循环
func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		res[0] = i
		for j := i + 1; j < len(nums); j++ {
			res[1] = j
			if nums[i]+nums[j] == target {
				return res
			}
		}
	}
	return nil
}

func twoSumOpt(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		mp[v] = i
	}
	for i, v := range nums {
		// 若是为[3,3]之类的条件，那么此次mp[3]=1,i=0,因此i必不会等于k
		if k, ok := mp[target-v]; ok && i != k {
			return []int{i, k}
		}
	}
	return nil
}
