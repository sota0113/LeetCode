package main

import (
	"fmt"
)

func main () {
	// input "["h","e","l","o","o"]
	// output "["o","o","l","e","h"]
	// must solve recursively.
	input := []byte("hello")
	reverseString(input)
}

func reverseString(s []byte) {
    if len(s) == 0 {
        return
    }
    left := 0
    right := len(s)-1
    ans := helper(left,right,s)
    fmt.Println(string(ans))
}

func helper(left,right int,s []byte ) []byte {
    lTmp := s[left]
    rTmp := s[right]
    s[right] = lTmp
    s[left] = rTmp
    left = left+1
    right = right-1
    if left < right {
        helper(left,right,s)
    }
    return s
}
