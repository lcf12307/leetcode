package main

/**
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5


*/
/**
 [-10,-3,0,5,9]
输出：[5,-3,9,-10,null,5]
预期：[0,-3,9,-10,null,5]
标准输出：[-10 -3 0 5 9]
*/
func run_5() {
	data := []int{-10, -3, 0, 5, 9}
	PrintTree(sortedArrayToBST(data))
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			nums[0],
			nil,
			nil,
		}
	}
	if len(nums) == 2 {
		return &TreeNode{
			nums[1],
			&TreeNode{
				nums[0],
				nil,
				nil,
			},
			nil,
		}
	}
	center := len(nums) / 2
	temp1 := nums
	temp2 := nums
	left := temp1[:center]
	right := temp2[center+1:]
	return &TreeNode{
		nums[center],
		sortedArrayToBST(left),
		sortedArrayToBST(right),
	}
}
