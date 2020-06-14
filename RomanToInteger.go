package main

import (
	"fmt"
	//"strings"
)

func main () {
	input := []string{"III","IV","IX","LVIII","MCMXCIV","MDCXCV","D"}
	for _,val := range input {
		fmt.Println("INPUT:",val)
		ans := romanToInt(val)
		fmt.Println("ANSWER :",ans)
	}
}

var m = map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
func romanToInt(s string) int {
	//map each Symbol to integer
/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
	result := 0
	skip := len(s)-1
	for i,_ := range s {	//if 7char -> 0-6
		if i == skip {
			continue
		}
		if i < len(s)-1 {	//len(s)-1 -> 6
			if m[string(s[i])] >= m[string(s[i+1])] {
				if i == len(s)-2 {
					result += m[string(s[i])]+m[string(s[i+1])]
				} else {
					result += m[string(s[i])]
				}
			} else {
				result += m[string(s[i+1])]-m[string(s[i])]
				skip = i+1
			}
		}
	}
        if skip == len(s)-2 || skip == 0 {
                result += m[string(s[len(s)-1])]
        }
	return result
}
