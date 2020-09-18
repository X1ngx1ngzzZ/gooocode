package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)


func main(){
	//获取输入
	reader:=bufio.NewReader(os.Stdin)
	a,_:=reader.ReadString('\n')
	input:=strings.Split(a,";")
	//fmt.Println(input[0],input[1])
	row:=strings.Fields(input[0])
	last:=strings.Fields(input[1])
	//把输入转成两个字符串切片
	tmp1:=string(row[len(row)-1][len(row[len(row)-1])-1])
	row[len(row)-1]=row[len(row)-1][:len(row[len(row)-1])-1]
	row=append(row,tmp1)
	tmp2:=string(last[len(last)-1][len(last[len(last)-1])-1])
	last[len(last)-1]=last[len(last)-1][:len(last[len(last)-1])-1]
	last=append(last,tmp2)
	//以last的长度作为总长度
	all:=len(last)
	res:=0
	i:=0
	j:=0

	//fmt.Println(row)
	//fmt.Println(last)
	// 以last为标准，去row里面每个元素看一样不一样
	for i=0;i<all;i++{
		//标志位初始化为否
		tmp:=false
		rowindex:=-1
		lastindex:=-1
		for j=0;j<len(row);j++{
			// 找到一样的了，标志位就置为真
			if last[i]==row[j]{
				tmp=true
				lastindex=i
				rowindex=j
				//fmt.Println(last[i],"find",row[j])
			}
			// 如果last当前的是这三个符号，row里面如果也有这三个符号就置为真
			if last[i]==","||last[i]=="？"||last[i]=="!"{
				if row[j]==","||row[j]==","||row[j]==","{
					tmp=true
				}
			}
		}
		// 每个元素找完一遍后，看标志位
		if tmp==true{
			//如果是对应元素，那么不用动直接下一轮循环
			if rowindex==lastindex{
				continue
			}else {
				res++
				//fmt.Println(i,j)
				//fmt.Println(res)
			}
		// 没找到，结果+1
		}else{
			res++
			//fmt.Println(res)
		}
	}
	s:=fmt.Sprintf("(%d,%d)",res-1,all)
	fmt.Println(s)
}