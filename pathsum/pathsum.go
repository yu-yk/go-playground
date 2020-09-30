package main

import "fmt"

func main() {
	n8 := TreeNode{8, nil, nil}
	n4 := TreeNode{4, nil, nil}
	n5 := TreeNode{5, &n4, &n8}

	test := pathSum(&n5, 9)
	fmt.Println(test)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	pathStack := []int{}

	dfs(root, sum, &pathStack, &res)
	return res
}

func dfs(root *TreeNode, target int, pathStack *[]int, res *[][]int) {
	if root == nil {
		*pathStack = append(*pathStack, 0)
		return
	}

	*pathStack = append(*pathStack, root.Val)

	if root.Left == nil && root.Right == nil && root.Val == target {
		path := make([]int, len(*pathStack))
		copy(path, *pathStack)
		*res = append(*res, path)
		return
	}

	dfs(root.Left, target-root.Val, pathStack, res)
	*pathStack = (*pathStack)[:len(*pathStack)-1]

	dfs(root.Right, target-root.Val, pathStack, res)
	*pathStack = (*pathStack)[:len(*pathStack)-1]
}
