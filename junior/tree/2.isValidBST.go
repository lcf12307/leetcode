package main

import "fmt"

/**
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

    节点的左子树只包含小于当前节点的数。
    节点的右子树只包含大于当前节点的数。
    所有左子树和右子树自身必须也是二叉搜索树。

示例 1:

输入:
    2
   / \
  1   3
输出: true

示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。


*/
func run_98() {
	//root := &TreeNode{
	//	3,
	//	&TreeNode{
	//		9,
	//		nil,
	//		nil,
	//	},
	//	&TreeNode{
	//		20,
	//		&TreeNode{
	//			15,
	//			nil,
	//			nil,
	//		},
	//		&TreeNode{
	//			7,
	//			nil,
	//			nil,
	//		},
	//	},
	//}
	root := &TreeNode{
		1,
		&TreeNode{
			1,
			nil,
			nil,
		},
		nil,
	}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	const INT_MAX = 1 << 31
	return judge(root, -INT_MAX-1, INT_MAX)
}

func judge(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}

	return judge(root.Left, min, root.Val) && judge(root.Right, root.Val, max)
}
