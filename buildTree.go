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
	//sample binary tree.
	/*
	    3
	   / \
	  9  20
	    /  \
	   15   7
	9
	 \
	  15
	 / \
       3    20
	     \
	      7
        9
       / \
      5   15
     /     \
    9      20
     \       \
      10      7
	*/
	inorder := []int{9,3,15,20,7} //left most
	postorder := []int{9,15,7,20,3} //Down most
	root := buildTree(inorder,postorder)
	fmt.Println(root)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return root
	}
	root := &TreeNode{postorder[len(postorder)-1],nil,nil}
	helper(root,inorder,postorder,index)
	return root
}

func helper(node *TreeNode,inorder []int, postorder []int,index int) {
	switch {
	case index == 0 :
		if inorder[index] == postorder[index] {
			node = &{inorder[index],nil,nil}
		} else {
			node = &{postorder[index],nil,nil}
		}
	//	index = index+1
	case index < len(inorder) :
		
	root.Val = inorder[index]
}
