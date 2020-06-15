package main

import (
	"fmt"
)


func main () {
	inputs := [][]string{{"a"}, {"aca","cba"},{"flower","flow","flw"},{"dog","racecar"}}
	//inputs := [][]string{{"a"}}
	//inputs := [][]string{{"aca","cba"}}
	for _,input := range inputs {
		result :=longestCommonPrefix(input)
		fmt.Println(result)
	}
}

func longestCommonPrefix(strs []string) string {
	result := ""
	if len(strs) == 0 {
		return result
	}
	if len(strs) == 1 {
		result = strs[0]
		return result
	}
	for i, char := range strs[0] {
		count := 0
		for num,val := range strs {
			if num == 0 || len(val) <= i {
				continue
			}
			if string(char) == string(val[i]) {
				count = count+1
			} else {
				goto LABEL
			}
			if count == len(strs)-1 {
				result += string(char)
			}
			continue
		}
	}
	LABEL:
	return result
}
