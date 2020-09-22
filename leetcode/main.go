package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 二分查找
func binarySearch(array *[]int, tar int) {
	h := len(*array) - 1
	l := 0
	if (*array)[l] > tar || (*array)[h] < tar {
		fmt.Println("target is not in range")
		return
	}
	for l <= h {
		if (*array)[(l+h)/2] == tar {
			fmt.Println("the index is: ", (l+h)/2)
			return
		} else if (*array)[(l+h)/2] < tar {
			l = (l + h) / 2
		} else {
			h = (l + h) / 2
		}
	}
	fmt.Println("target is not in array")
}

// two sum
func twoSum(array []int, tar int) []int {
	target := make(map[int]int)
	for k, v := range array {
		_, ok := target[v]
		if ok {
			return []int{target[v], k}
		} else {
			target[tar-v] = k
		}
	}
	return nil
}

// add two nums
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num int
	tar := &ListNode{0, nil}
	head := tar
	for {
		if l1 != nil || l2 != nil {
			num /= 10
			if l1 != nil {
				num += l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				num += l2.Val
				l2 = l2.Next
			}
			head.Next = &ListNode{Val: num % 10}
			head = head.Next
		}
		if num/10 != 0 {
			head.Next = &ListNode{Val: 1}
			head = head.Next
		}
		return tar.Next

	}
}

//
func main() {
	//var a = []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	//a = a[:]
	//var b = 13
	//binarySearch(&a, b)

	//var a = []int{1,4,7,9}
	//var b = 10
	//fmt.Println(twoSum(a, b))

	//n1 := &ListNode{
	//	3,
	//	nil,
	//}
	//n2 := &ListNode{
	//	4,
	//	n1,
	//}
	//n3 := &ListNode{
	//	2,
	//	n2,
	//}
	//n4 := &ListNode{
	//	4,
	//	nil,
	//}
	//n5 := &ListNode{
	//	6,
	//	n4,
	//}
	//n6 := &ListNode{
	//	5,
	//	n5,
	//}
	//fmt.Println(addTwoNumbers(n3, n6).Val)

}
