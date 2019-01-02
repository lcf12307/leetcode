package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(ln *ListNode) {
	if ln == nil {
		fmt.Println("==================")
		return
	}

	fmt.Println(ln.Val)
	PrintListNode(ln.Next)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTree(root *TreeNode){
	if root == nil {
		fmt.Print("null, ")
	} else {
		fmt.Print(root.Val, ", ")
		PrintTree(root.Left)
		PrintTree(root.Right)
	}
}