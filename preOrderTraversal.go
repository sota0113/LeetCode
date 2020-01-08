package main

import (
	"fmt"
)
// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main () {
	// initialize sample tree.
	three := TreeNode{3, nil, nil}
	two := TreeNode{2, nil, &three}
	one := TreeNode{1, nil, &two}
	result := preorderTraversal(&one)
	fmt.Println(result)
}

func preorderTraversal(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }

    helper(&res, root)
    return res
}

func helper(res *[]int, root *TreeNode) {
    *res = append(*res, root.Val)
    if root.Left != nil {
        helper(res,root.Left)
    }
    if root.Right != nil {
        helper(res,root.Right)
    }
}
