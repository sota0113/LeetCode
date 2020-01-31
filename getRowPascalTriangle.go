//パスカルの三角形
package main

import (
	"fmt"
)

func main () {
//	**** smple 1 ****
//      Input: 0
//      Output:
//      [
//      ]

//	**** smple 2 ****
//      Input: 3
//      Output:
//      [
//         [1,2,1]
//      ]

//	**** smple 3 ****
//      Input: 5

//       [1,4,6,4,1]
//      ]
	//ans := generate(6)
	ans := getRow(0)
	fmt.Println(ans)
}

func getRow(rowIndex int) []int {
	fstLine := []int{1}
	if rowIndex < 1 {
		return fstLine
	}
	ans := helper(fstLine,rowIndex)
	return ans
}

func helper(fstLine []int, rowIndex int ) []int {
	var secLine []int
	secLine = append(secLine,1)
	for i := 0 ; i<len(fstLine)-1 ;i++ {
		secLine = append(secLine,fstLine[i]+fstLine[i+1])
	}
	secLine = append(secLine,1)
	if rowIndex != len(fstLine) {
		secLine = helper(secLine,rowIndex)
	}
	return secLine
}
