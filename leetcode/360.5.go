package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	if len(s)>3000{
		return
	}
	slice:=strings.Split(s,"n")
	for i:=0;i<len(slice);i++{
		if i==0{
			//fmt.Println(slice[i])
			tmp:=strings.ToUpper(slice[i][:1])
			//fmt.Println(tmp)
			slice[i]=slice[i][1:]
			slice[i]=tmp+slice[i]

		}
		if i>0{
			slice[i]="N"+slice[i]
		}
	}
	for _,v:=range slice{
		fmt.Println(v)
	}

}
