package main

import (
	"fmt"

	"github.com/baqimeilaoban/leetcode/simple"
)

func main() {
	strs := []string{"leets", "leetcode", "lee", "leets"}
	res := simple.LongestCommonPrefixOpt(strs)
	fmt.Println(res)
}
