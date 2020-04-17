package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
 }

func main () {
	//initialization
	// 5 -> 4 -> 3 -> 2 -> 1 -> nil
	// expected output
	// 1 -> 2 -> 3 -> 4 -> 5 -> nil
	head := &ListNode{5,nil}
	head.Next = &ListNode{4,nil}
	head.Next.Next = &ListNode{3,nil}
	head.Next.Next.Next = &ListNode{2,nil}
	head.Next.Next.Next.Next = &ListNode{1,nil}
	ans := reverseList(head)
	display(ans)
}

func display (ans *ListNode) {
	if ans == nil {
		fmt.Println("ERROR")
	}
	if ans.Next != nil {
		fmt.Println(ans.Val)
		fmt.Println(ans.Next)
		display(ans.Next)
	}
}

func reverseList(head *ListNode) *ListNode {
			//head
	// 5 -> 4 -> 3 -> 2 -> 1 -> nil
	prev := &ListNode{}
	newLine := &Listnode{}
	helper(prev,head,newLine)
	return head
}

func helper(prev, current, newLine *ListNode) {
	fmt.Println("top current:",current.Val)
	if current.Next != nil {
		helper(current,current.Next)
	}
	newLine.Val = current.Val
	newLine.Next = prev
}
