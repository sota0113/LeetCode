package main

import (
	"fmt"
)

func main () {
	nums := []int{2,7,11,15}
	target := 26
	result := twoSum(nums,target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	ans := []int{}
	for i := 0 ; i<len(nums) ; i++ {
		for k := i+1 ; k < len(nums) ; k++ {
			//fmt.Printf("i=%v ,k=%v \n",i,k)
			if target == nums[i]+nums[k] {
				ans = []int{i,k}
				return ans
			}
		}
	}
	return ans
}
