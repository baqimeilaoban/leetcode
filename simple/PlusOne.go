package simple

// 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			if digits[0] == 9 && i == 0 {
				digits[0] = 1
				digits = append(digits, 0)
				return digits
			}
			digits[i] = 0
		}
	}
	return digits
}
