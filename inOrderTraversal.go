//中順巡回探索

package main

import (
	"fmt"
)

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main () {
	/*initialize sample tree
	1
	 \
	  2
	 /
	3
	*/
	one := &TreeNode{1,nil,nil}
	one.Right = &TreeNode{2,nil,nil}
	one.Right.Left = &TreeNode{3,nil,nil}
/*
	three := TreeNode{3,nil,nil}
	two := TreeNode{2,&three,nil}
	one := TreeNode{1,nil,&two}
	*/

	result :=inorderTraversal(one)
	fmt.Println(result)
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	helper(&res, root)
	return res
}

func helper (res *[]int, root *TreeNode) {
	if root.Left != nil {
		helper(res,root.Left)
	}
	*res = append(*res,root.Val)
	if root.Right != nil {
		helper(res,root.Right)
	}
}
