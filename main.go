package main

import (
	"fmt"

	"github.com/baqimeilaoban/leetcode/simple"
)

func main() {
	strs := []string{"dog", "racecar", "car"}
	res := simple.LongestCommonPrefix(strs)
	fmt.Println(res)
}
