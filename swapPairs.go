package main

import (
	"fmt"
)

//Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }

func main() {
	//init linked list as below.
	// 1 -> 2 -> 3 -> 4
	//expected answer as below.
	// 2 -> 1 -> 4 -> 3
	head := ListNode{1,nil}
	head.Next = &ListNode{2,nil}
	head.Next.Next = &ListNode{3,nil}
	head.Next.Next.Next = &ListNode{4,nil}

	//show head structure both before exec swapPairs function and after.
	fmt.Println("[ BEFORE exec swapPairs func ]")
	showLinkedList(&head)	//before
	ans := swapPairs(&head)
	fmt.Println("[ AFTER exec swapPairs func ]")
	showLinkedList(ans)	//after
}

func showLinkedList (head *ListNode) {
	if head == nil {
		fmt.Println("head is nil.")
	}
	if head.Next != nil {
		fmt.Printf(" %v ->",head.Val)
		showLinkedList(head.Next)
	} else {
		fmt.Printf(" %v\n",head.Val)
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	headVal := head.Val
	nextVal := head.Next.Val
	head.Val = nextVal
	head.Next.Val = headVal
	if head.Next.Next != nil {
		swapPairs(head.Next.Next)
	}
	return head
}
