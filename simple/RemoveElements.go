package simple

func RemoveElements(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[slow] = nums[i]
			slow++
		}
	}
	return slow
}
