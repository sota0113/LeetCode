package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main () {
	// initialization
	// 2 -> 4 -> 3 -> nil
	l1 := ListNode{2,nil}
	l1.Next = &ListNode{4,nil}
	l1.Next.Next = &ListNode{3,nil}
	// 5 -> 6 -> 4 -> nil
	l2 := ListNode{5,nil}
	l2.Next = &ListNode{6,nil}
	l2.Next.Next = &ListNode{4,nil}
	// expectation
	// 2->4->3->nil --[convert]--> 243 --[reverse]--> 342 
	// 5->6->4->nil --[convert]--> 564 --[reverse]--> 465
	// 342 + 465 = 807
	// 807 --[reverse]--> 708
	// 708 --[convert]--> 7->0->8->nil

	result := addTwoNumbers(&l1,&l2)
	depth,ans := depthSearch("",result,1)
	fmt.Println("DEPTH:",depth)
	fmt.Println("ANS:",ans)
//
}

func depthSearch (answer string,result *ListNode,depth int) (int,string) {
	var value int
	var ans string
	answer = answer+strconv.Itoa(result.Val)
	if result.Next !=nil {
		depth = depth+1
		value,ans = depthSearch(answer,result.Next,depth)
	} else {
		value = depth
		ans = answer
	}
	return value,ans
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1List, l2List string
	inputToString(l1,&l1List)
	inputToString(l2,&l2List)
	l1ListInt,_ := strconv.Atoi(l1List)
	l2ListInt,_ := strconv.Atoi(l2List)
	ansStrRvs := reverseString(strconv.Itoa(l1ListInt+l2ListInt))
	ansStrRvsArr := strings.Split(ansStrRvs,"")
	//recursrive
	result := strToListNode(ansStrRvsArr,nil)
	return result
}

func reverseString (input string) string {
	if len(input) == 0 {
		return input
	} else {
		var output string
		for i:=1 ; i<=len(input) ; i++ {
			output = output+string(input[len(input)-i])
		}
		return output
	}
}
//recursive
// 7(アドレス欲しい) -> 0 -> 8 -> nil
func strToListNode (inputArr []string, output *ListNode) *ListNode {
	var listNode ListNode
	var result *ListNode
	listNode.Val,_ = strconv.Atoi(inputArr[len(inputArr)-1])
	listNode.Next = output
	if len(inputArr[:len(inputArr)-1]) != 0 {
		result = strToListNode(inputArr[:len(inputArr)-1],&listNode)
	} else {
		result = &listNode
	}
	return result
}

func inputToString (list *ListNode, listString *string) {
	if list.Next != nil {
		inputToString(list.Next,listString)
	}
	*listString = *listString+strconv.Itoa(list.Val)
}
