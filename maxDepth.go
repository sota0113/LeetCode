//深さ探索

package main

import (
	"fmt"
	//"reflect"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val int
	Right *TreeNode
	Left *TreeNode
}

func main() {
	//Sample Tree structure.
	/*
	    3
	   / \
	  9  20
	    /  \
	   15   7
	  /
	 8
	  \
	   5
	*/
	//initialization.
	root := &TreeNode{3,nil,nil} //root is "*main.TreeNode" Type.
	root.Left = &TreeNode{9,nil,nil}
	root.Right = &TreeNode{20,nil,nil}
	root.Right.Left = &TreeNode{15,nil,nil}
	root.Right.Right = &TreeNode{7,nil,nil}
	root.Right.Left.Left = &TreeNode{8,nil,nil}
	root.Right.Left.Left.Right = &TreeNode{5,nil,nil}
	result := maxDepth(root)
	fmt.Println(result)
}


func maxDepth (root *TreeNode) int {
	if root == nil {
		return 0
	}
        if root.Left == nil && root.Right == nil {
                return 1
        }
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left+1
	} else {
		return right+1
	}
}
