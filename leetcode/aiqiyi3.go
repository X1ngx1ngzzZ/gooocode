package main

import "fmt"

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
	var arry  []byte
	m := map[byte]byte {
		')' : '(',
		']' : '[',
		'}' : '{',
	}
	for _, value := range s {
		clen := len(arry)
		if clen>0 {
			if _,ok := m[byte(value)]; ok {
				if arry[clen-1]==m[byte(value)] {
					arry = arry[:clen-1]
					continue
				}
			}
		}
		arry = append(arry,byte(value))
	}
	return len(arry)==0
}

