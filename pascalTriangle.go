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
//           [1],
//          [1,1],
//         [1,2,1]
//      ]

//	**** smple 3 ****
//      Input: 5
//      Output:
//      [
//           [1],
//          [1,1],
//         [1,2,1],
//        [1,3,3,1],
//       [1,4,6,4,1]
//      ]
	ans := generate(6)
	fmt.Println(ans)
}

func generate(numRows int) [][]int {
	var line [][]int
	if numRows < 1 {
		return line
	}
	ans := helper( &line,numRows,0)
	return *ans
}


func helper(line *[][]int,numRows,counter int) *[][]int {
	var tmpLine []int
	if counter == 0 {
		tmpLine = append(tmpLine,1)
		*line = append(*line,tmpLine)
	} else {
		// include conter+1 factors.
		//first and last factor is 1.
		//every other factors are sum of factor[n] and factor[n-1] of the previous line.
		prevLine := (*line)[counter-1]
		tmpLine = append(tmpLine,1)
		for i := 1 ; i < counter ;i++ { //counter factors are necessary and 2 factors,first and last, are not.
						//if the 'counter' is 3, loop twice and i=1 and 2.
			num := prevLine[i-1]+prevLine[i]
			tmpLine = append(tmpLine,num)
		}
		tmpLine = append(tmpLine,1)  //[1, ... ,1]
		*line = append(*line,tmpLine)
	}
	counter = counter+1
	if counter < numRows {
		helper(line,numRows,counter)
	}
	return line
}
