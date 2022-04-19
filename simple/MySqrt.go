package simple

import "math"

// 给你一个非负整数x，计算并返回x的算术平方根

func MySqrt(x int) int {
	if x == 0 {
		return 0
	}
	i := 1
	for true {
		if i*i == x {
			return i
		}
		if i*i < x && (i+1)*(i+1) > x {
			return i
		}
		i += 1
	}
	return 0
}

// MySqrtOpt 使用换底公式做
func MySqrtOpt(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}
