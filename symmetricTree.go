//対称木確認

package main

import(
	"fmt"
	"reflect"
	"strconv"
)

//definition of tree.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main () {
	//init tree.
	//sample tree structure no.1.
	//should return true.
	/*
	    1
	   / \
	  2   2
	 / \ / \
	3  4 4  3
	*/
	root := &TreeNode{1,nil,nil}
	root.Left = &TreeNode{2,nil,nil}
	root.Right = &TreeNode{2,nil,nil}
	root.Left.Left = &TreeNode{3,nil,nil}
	root.Left.Right = &TreeNode{4,nil,nil}
	root.Right.Left = &TreeNode{4,nil,nil}
	root.Right.Right = &TreeNode{3,nil,nil}
	//sample tree structure no.2.
	//should return false.
	/*
	    1
	   / \
	  2   2
	   \   \
	   3    3
	*/
	/*
	root = &TreeNode{1,nil,nil}
	root.Left = &TreeNode{2,nil,nil}
	root.Left.Right = &TreeNode{3,nil,nil}
	root.Right = &TreeNode{2,nil,nil}
	root.Right.Right = &TreeNode{3,nil,nil}
	*/
	result := isSymmetric(root)
	fmt.Println(result)
}

func isSymmetric (root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := []string{}
	right := []string{}
	leftPriorSearch(root,&left)
	rightPriorSearch(root,&right)
	fmt.Println(left)
	fmt.Println(right)
	return reflect.DeepEqual(left,right)
}

func leftPriorSearch (root *TreeNode,leftNode *[]string) {
	//add root.Val to leftNode
	if root != nil {
		if root.Left != nil {
			*leftNode = append(*leftNode,strconv.Itoa(root.Left.Val))
		        leftPriorSearch(root.Left,leftNode)
		} else {
			*leftNode = append(*leftNode,"nil")
		}
		if root.Right != nil {
			*leftNode = append(*leftNode,strconv.Itoa(root.Right.Val))
		        leftPriorSearch(root.Right,leftNode)
		} else {
			*leftNode = append(*leftNode,"nil")
		}
	}
}

func rightPriorSearch (root *TreeNode,rightNode *[]string) {
	//add root.Val to rightNode
	if root != nil {
		if root.Right != nil {
			*rightNode = append(*rightNode,strconv.Itoa(root.Right.Val))
		        rightPriorSearch(root.Right,rightNode)
		} else {
			*rightNode = append(*rightNode,"nil")
		}
		if root.Left != nil {
			*rightNode = append(*rightNode,strconv.Itoa(root.Left.Val))
			rightPriorSearch(root.Left,rightNode)
		} else {
			*rightNode = append(*rightNode,"nil")
		}
	}
}
