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
	/* tree structure sample 1.
	    3
	   / \
	  9  20
	    /  \
	   15   7
	*/
        /*root := &TreeNode{3,nil,nil}
        root.Left = &TreeNode{9,nil,nil}
        root.Right = &TreeNode{20,nil,nil}
        root.Right.Left = &TreeNode{15,nil,nil}
        root.Right.Right = &TreeNode{7,nil,nil}
        root := &TreeNode{3,nil,nil}
	*/
	/* tree structure sample 2.
	    1
	   / \
	  2   3
	 / \
	4   5
	*/
        root := &TreeNode{1,nil,nil}
        root.Left = &TreeNode{2,nil,nil}
        root.Right = &TreeNode{3,nil,nil}
        root.Left.Left = &TreeNode{4,nil,nil}
        root.Left.Right = &TreeNode{5,nil,nil}
	result := levelOrder(root)
	fmt.Println(result)
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	depth := -1
	if root == nil {
		return res
	}
	helper(depth,&res,root)
	return res
}

func helper(depth int,res *[][]int, root *TreeNode) {
	depth = depth+1 // start from 0
	if len(*res)-1 < depth { // len(*res) is start from 0.
		init := []int{}
		*res = append(*res, init)
	}
	(*res)[depth] = append((*res)[depth], root.Val)
	if root.Left != nil {
		helper(depth,res,root.Left)
	}
	if root.Right != nil {
		helper(depth,res,root.Right)
	}
}
