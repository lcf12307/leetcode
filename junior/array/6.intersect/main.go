package main

import (
	"fmt"
)

/*

给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]
说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
我们可以不考虑输出结果的顺序。
进阶:

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

*/
//解法1：直接on2的方法，进行比对;
//解法2：利用异或可以获得较大集合关于两集合交集的补集，再求一次补集
//解法3：利用map来计算
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		nums2, nums1 = nums1, nums2
	}
	temp := make(map[int]int)
	var result []int
	for _, data := range nums1 {
		if _, ok := temp[data]; ok {
			temp[data]++
		} else {
			temp[data] = 1
		}
	}
	for _, data := range nums2 {
		if _, ok := temp[data]; ok && temp[data] > 0 {
			result = append(result, data)
			temp[data]--
		}
	}
	return result
}

func main() {

	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	result := intersect(nums1, nums2)
	fmt.Println(result)
}
