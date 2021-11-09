package leetcode

/**
101. 对称二叉树
给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

*/
func isSymmetric(root *TreeNode) bool {

	return check(root, root)
}

func check(p, q *TreeNode) bool{
	if p == nil && q == nil{
		return true
	}

	if p == nil || q == nil{
		return false
	}

	return p.Val == q.Val && check(p.Left, q.Right) && check(q.Left, p.Right)
}