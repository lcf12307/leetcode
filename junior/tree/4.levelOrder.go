package main

import "fmt"

/**

给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]


 */
func run_4()  {
	root := &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{
				3,
				&TreeNode{
					4,
					&TreeNode{
						5,
						nil,
						nil,
					},
					nil,
				},
				nil,
			},
			nil,
		},
		nil,
	}

	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var column [] int
	if root == nil {
		return result
	}
	var stack []*TreeNode
	current := 1
	next := 0
	stack = append(stack, root)
	for len(stack) > 0  {
		temp := stack[0]
		stack = stack[1:]
		current --
		if temp != nil {
			stack = append(stack, temp.Left, temp.Right)
			next += 2
			column = append(column, temp.Val)
		}
		if current == 0 {
			current, next = next, 0
			if len(column) > 0 {
				result = append(result, column)
			}
			column = []int {}
		}
	}
	if len(column) > 0 {
		result = append(result, column)
	}
	return result
}