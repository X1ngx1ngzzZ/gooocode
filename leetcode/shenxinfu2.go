package main

import (
	"fmt"
	"sort"
)

func commonChars( chars []string ) string {
	// write code here
	s:=make([]int,0)
	res:=""
	m:=map[int]int{}
	for i:=0;i<len(chars);i++{
		for _,v:=range chars[i]{
			m[int(v)]++
		}
	}

	for k,v:=range m{
		for v>=len(chars){
			s=append(s,k)
			v-=len(chars)
		}
	}
	sort.Ints(s)
	for i:=0;i<len(s);i++{
		res+=string(s[i])
	}
	return res

}

func main(){
	arry:=[]string{"bella","label","roller"}
	fmt.Println(commonChars(arry))
}

