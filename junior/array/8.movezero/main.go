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
func moveZeroes(nums []int) {
	//记录0的个数
	count := 0
	for index := 0; index < len(nums); index++ {
		if nums[index] == 0 {
			count++
			if count+index > len(nums) {
				break
			}
			nums = append(nums[:index], nums[index+1:]...)
			nums = append(nums, 0)
			index--
		}
	}
}

func main() {

	nums1 := []int{0, 0, 1}
	moveZeroes(nums1)
	fmt.Println(nums1)
}
