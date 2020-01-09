package main

import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main () {
/*
   1
    \
     2
    /
   3
*/
	three := TreeNode{3,nil,nil}
	two := TreeNode{2,&three,nil}
	one := TreeNode{1,nil,&two}

	result := postorderTraversal(&one)
	fmt.Println(result)
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	helper (&res,root)
	return res
}


func helper (res *[]int, root *TreeNode) {
	if root.Left != nil {
		helper(res,root.Left)
	}
	if root.Right != nil {
		helper(res,root.Right)
	}
	*res = append(*res, root.Val)
}
