package simple

// 两数之和

func TwoSum(nums []int, target int) []int {
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

func TwoSumOpt(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, num := range nums {
		if p, ok := hashTable[target-num]; ok {
			return []int{p, i}
		}
		hashTable[num] = i
	}
	return nil
}
