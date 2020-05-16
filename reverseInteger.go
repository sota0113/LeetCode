package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main () {
	//sample inputs
	input := [...]int{123, -123, 120, 1534236469}
	for _,val := range input {
		//fmt.Printf("num: %v, val: %v\n",num,val)
		result := reverse (val)
		fmt.Println(result)
	}
}
func reverse(x int) int {
	if x < -2147483648 || x > 2147483647 {
		x = 0
	}
	isMinus := false
	if x < 0 {
		isMinus = true
	}
	abs := math.Abs(float64(x))
	val := strconv.Itoa(int(abs))
	varArr := []string{}
	for _, k := range val {
		varArr = append([]string{string(k)},varArr[0:]...)
	}
	varInt,_ := strconv.Atoi(strings.Join(varArr,""))
	if isMinus == true {
		varInt = -varInt
	}
        if varInt < -2147483648 || varInt > 2147483647 {
                varInt = 0
        }
	return int(varInt)
}
