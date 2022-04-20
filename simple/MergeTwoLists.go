package simple

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 边界条件
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists(list2.Next, list1)
		return list2
	}
}
