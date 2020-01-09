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
	three := TreeNode{3,nil,nil}
	two := TreeNode{2,&three,nil}
	one := TreeNode{1,nil,&two}
	result :=inorderTraversal(&one)
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
