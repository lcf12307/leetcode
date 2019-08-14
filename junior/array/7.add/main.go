package main

import (
	"fmt"
)

/*

给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/
func plusOne(digits []int) []int {

	length := len(digits)

	if digits[length-1] == 9 {
		digits = plus(length-1, digits)
	} else {
		digits[length-1]++
	}
	return digits
}

func plus(index int, digits []int) []int {
	if index == 0 && digits[0] == 9 {
		var temp = []int{0}
		digits = append(temp, digits...)
		index = 1
	}
	digits[index] = 0
	if digits[index-1] == 9 {
		digits = plus(index-1, digits)
	} else {
		digits[index-1]++
	}
	return digits
}
func main() {

	nums1 := []int{9, 9}
	result := plusOne(nums1)
	fmt.Println(result)
}
