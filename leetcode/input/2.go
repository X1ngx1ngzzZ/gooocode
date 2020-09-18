package main

import (
	"bufio"
	"fmt"
	"strings"

	"os"
)

func main(){
	var a int
	fmt.Scan(&a)
	reader:=bufio.NewReader(os.Stdin)
	input:=make([]string,a)
	var result [][]string
	for i:=0;i<a;i++{
		input[i],_=reader.ReadString('\n')
		input[i]=strings.TrimSpace(input[i])
		tmp:=strings.Fields(input[i])
		result=append(result,tmp)
	}

	var res int
	m:=make(map[string]int)
	for i:=0;i<len(result);i++{
		for j:=0;j<len(result[i]);j++{
			if j==0{
				m[result[i][j]]++
			}
			if j==1{
				if _,ok:=m[result[i][j]];ok{
					res++
					delete(m,result[i][j])
				}
			}
		}
	}
	fmt.Println(res)
}