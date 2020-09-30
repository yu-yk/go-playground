package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	//bfs
	sum := 0
	currentLevel := []*TreeNode{root}

	for len(currentLevel) > 0 {
		nextLevel := []*TreeNode{}
		sum = 0
		for _, v := range currentLevel {
			sum += v.Val
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}
		}
		currentLevel = nextLevel
	}
	return sum
}
