package main

import (
	"fmt"

	"github.com/baqimeilaoban/leetcode/simple"
)

func main() {
	list13 := &simple.ListNode{
		Val:  4,
		Next: nil,
	}
	list12 := &simple.ListNode{
		Val:  2,
		Next: list13,
	}
	list11 := &simple.ListNode{
		Val:  1,
		Next: list12,
	}

	list23 := &simple.ListNode{
		Val:  4,
		Next: nil,
	}
	list22 := &simple.ListNode{
		Val:  3,
		Next: list23,
	}
	list21 := &simple.ListNode{
		Val:  1,
		Next: list22,
	}
	list := simple.MergeTwoLists(list11, list21)
	fmt.Println(list.Val, list.Next.Val)
}
