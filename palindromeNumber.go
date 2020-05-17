package main

import (
	"fmt"
)

func main () {
	imput := []int{121,-121,10,56789,1234321}
	for _,val := range imput {
		fmt.Printf("VAL:%v RESULT:%v \n",val,isPalindrome(val))
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	list := []int{}
	list = intToSlice(x,list)

	for i,_ := range list {
		k := len(list)-1-i
		if i <= k {
			if list[i] != list[len(list)-1-i] {
				return false
			}
		}
	}
	return true
}

func intToSlice (x int, list []int) []int {
//	fmt.Println("X:",x)
//	fmt.Println("list:",list)
//	fmt.Println("x % 10:",x%10)
//	fmt.Println("x / 10:",x/10)
	if x > 0 {
		return intToSlice(x/10,append([]int{x%10},list[0:]...))
	}
	return list
}
