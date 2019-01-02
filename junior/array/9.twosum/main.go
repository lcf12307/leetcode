package main

import (
	"fmt"
)

/*

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
 */

func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	for i,d := range nums{
		if _, ok := temp[target - d]; ok {
			return []int {temp[target - d], i}
		}
		temp[d] = i
	}
	return []int {}
}

func main()  {


	nums1 := []int{2,7,11,15}
	result := twoSum(nums1,9)
	fmt.Println(result)
}

