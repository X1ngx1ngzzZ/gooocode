package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	res:=deal(s)
	if res==true{
		fmt.Println("True")
	}else {
		fmt.Println("False")
	}

}

func deal(s string) bool {
	x,y := 0,0
	m := make(map[string]int,10)
	m["00"] = 1
	var ok bool
	tmp := ""
	for i:=0;i<len(s);i++{
		if s[i] == 'N'{
			y += 1
		}else if s[i] == 'S'{
			y-=1
		}else if s[i] == 'W'{
			x -= 1
		}else{
			x += 1
		}
		tmp = strconv.Itoa(x) + strconv.Itoa(y)
		_,ok = m[tmp]
		if !ok{
			m[tmp] = 1
		}else{
			return true
		}
	}
	return false
}


