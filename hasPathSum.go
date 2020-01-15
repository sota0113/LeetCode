package main

import (
	"fmt"
	//"reflect"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main () {
	//sample tree structure.
	/*
	      5
	     / \
	    4   8
	   /   / \
	  11  13  4
	 /  \      \
	7    2      1
	*/
	//initialization.
	/*root := &TreeNode{5,nil,nil}
	root.Left = &TreeNode{4,nil,nil}
	root.Right = &TreeNode{8,nil,nil}
	root.Left.Left = &TreeNode{11,nil,nil}
	root.Left.Left.Left = &TreeNode{7,nil,nil}
	root.Left.Left.Right = &TreeNode{2,nil,nil}
	root.Right = &TreeNode{8,nil,nil}
	root.Right.Left = &TreeNode{13,nil,nil}
	root.Right.Right = &TreeNode{4,nil,nil}
	root.Right.Right.Right = &TreeNode{1,nil,nil}

	sum := 22
	*/
	//sample tree structure.
        /*
              1
             / \
            -2  -3
           /   / \
          1   3  -2
           \
           -1
        */
        //initialization.
        root := &TreeNode{1,nil,nil}
	root.Left = &TreeNode{-2,nil,nil}
	root.Right = &TreeNode{-3,nil,nil}
	root.Left.Left = &TreeNode{1,nil,nil}
	root.Right.Left = &TreeNode{3,nil,nil}
	root.Right.Right = &TreeNode{-2,nil,nil}
	root.Left.Left.Right = &TreeNode{-1,nil,nil}
	sum := 3
	result := hasPathSum(root,sum)
	fmt.Println(result)
}

//Left Prior Search
func hasPathSum(root *TreeNode, sum int) bool {
	tmp := 0
	result := false
	if root == nil {
		return result
	}
	helper(root,&tmp,&result,sum)
	return result
}

func helper(root *TreeNode, tmp *int, result *bool, sum int) {
        *tmp += root.Val
        if root.Left != nil {
                helper(root.Left,tmp,result,sum)
        }
        if root.Right != nil {
                helper(root.Right,tmp,result,sum)
        }
        if root.Right == nil && root.Left == nil {
                fmt.Println("[TMP]",*tmp)
                fmt.Println("[SUM]",sum)
                if *tmp == sum {
                        *result = true
                }
        }
	*tmp -= root.Val
}
