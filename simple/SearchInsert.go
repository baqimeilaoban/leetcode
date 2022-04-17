package simple

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] < target {
			return 1
		}
		return 0
	}
	// 针对 nums 数量大于2
	for i := 0; i <= len(nums)-1; i++ {
		if target == nums[i] {
			return i
		} else if nums[i] < target && nums[i+1] > target {
			return i + 1
		} else if target > nums[len(nums)-1] {
			return len(nums)
		}
	}
	return 0
}

// SearchInsertOpt 使用二分法优化
func SearchInsertOpt(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
