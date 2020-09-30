package binarytree

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	inorderDFS(root, &res)
	fmt.Println(res)
	return res
}

func inorderDFS(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	inorderDFS(root.Left, res)
	*res = append(*res, root.Val)
	inorderDFS(root.Right, res)
}

func inorderTraversal2(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr) // push all left node
			curr = curr.Left
		}

		curr = stack[len(stack)-1] // no more left, pop stack top
		stack = stack[:len(stack)-1]

		res = append(res, curr.Val)
		curr = curr.Right
	}

	return res
}
