package leetcode

/*
https://leetcode-cn.com/problems/invert-binary-tree/
翻转二叉树
	翻转一棵二叉树。
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left
	return root
}