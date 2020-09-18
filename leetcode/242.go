package main

import (
	"fmt"
	"sort"
	"strings"
)
func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	slice1:=strings.Split(s,"")
	slice2:=strings.Split(t,"")
	sort.Strings(slice1)
	//fmt.Println(slice1)
	sort.Strings(slice2)
	//fmt.Println(slice2)
	for i:=0;i<len(s);i++{
		if slice2[i]!=slice1[i]{
			//fmt.Println(s[i],t[i])
			return false
		}
	}
	return true
}

func main(){
	fmt.Println(isAnagram("anagram","nagaram"))
}
