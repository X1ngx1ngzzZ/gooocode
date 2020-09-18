package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//sum:=0
	reader:=bufio.NewReader(os.Stdin)

	//input,_:=reader.ReadString('\n')
	//去空格
	//input=strings.TrimSpace(input)
	//fmt.Println(input)
	//result:=strings.Fields(input)
	//fmt.Println(result)
	//for _,v:=range result{
	//	tmp,_:=strconv.Atoi(v)
	//	fmt.Println(tmp)
	//	sum+=tmp
	//可以使用空接口接收
	input:=make([]interface{},4)
	for i:=0;i<=3;i++ {
	input[i],_=reader.ReadString('\n')
	}
	for _,v:=range input{
		fmt.Println(v)

	}

	//fmt.Println(sum)
}