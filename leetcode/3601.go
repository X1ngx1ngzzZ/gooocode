package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main(){
	reader:=bufio.NewReader(os.Stdin)
	var num int
	fmt.Scan(&num)
	if num>2000&&num<1{
		fmt.Println(-1)
	}
	input:=make([]string,num)
	var res int

	for i:=0;i<len(input);i++ {
		input[i],_=reader.ReadString('\n')
	}
	for _,v:=range input{
		v=strings.TrimSpace(v)
	}
	for k,v:=range input{
			if len(v)<=10 {
				matched, _ := regexp.MatchString("^[A-Za-z]+$", v)
				if matched == true {
					fmt.Println(k,"is true")
					res++
				}
			}
	}
	fmt.Println(res)

}
