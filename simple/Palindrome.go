package simple

// 回文数

func IsPalindrome(x int) bool {
	// 特殊情况
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	tempNum := 0
	for x > tempNum {
		tempNum = tempNum*10 + x%10
		x = x / 10
	}
	return x == tempNum || x == tempNum/10
}
