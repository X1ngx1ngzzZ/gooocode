package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var m,n int
	fmt.Scan(&n,&m)
	//fmt.Println(m,n)
	input:=make([]string,m)
	reader:=bufio.NewReader(os.Stdin)
	for i:=0;i<m;i++ {
		tmp,_:=reader.ReadString('\n')
		input[i]=strings.TrimSpace(tmp)
	}
	//fmt.Println("shuzu:",input)
	res:=0

	for i:=1;i<=n;i++{
		flag:=false
		//fmt.Println(i)
		for j:=0;j<len(input);j++{
			tmp,_:=strconv.Atoi(input[j])
			//fmt.Println(tmp)
			if i%tmp==0{
				flag=true
			}
		}
		if flag==true{
			res++
		}
	}
	fmt.Println(res)
}
