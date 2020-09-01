package main

import "fmt"

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
		} else if (*array)[l+h]/2 < tar {
			h = (l + h) / 2
		} else {
			l = (l + h) / 2
		}
		fmt.Println("target is not in array")
		return
	}
}

func main() {
	var a = []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	a = a[:]
	var b = 9
	binarySearch(&a, b)
}
