package main

import (
	"fmt"
	sort2 "sort"
)

func main() {
	var s string
	fmt.Scan(&s)
	var k int
	fmt.Scan(&k)
	slice:=getAllSubstring(s)
	input:=RemoveRepeatedElement(slice)
	sort2.Strings(input)
	fmt.Println(input)
	res:=deal(input,k)
	fmt.Println(res)
}

func getAllSubstring(str string) (re []string) {
	for i := 0; i < len(str); i++ {
		j := len(str)
		for j = len(str); j > i; j-- {
			re = append(re, str[i:j])
		}
	}
	return re
}
func deal(arry []string,k int)string{

	return arry[k-1]

}
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}