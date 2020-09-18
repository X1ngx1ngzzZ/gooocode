package main

import (
	"bufio"
	//"flag"
	"fmt"
	"os"
	"strconv"

	//"strconv"
	"strings"
)
//
//func main(){
//	var a int
//	reader:=bufio.NewReader(os.Stdin)
//	fmt.Scan(&a)
//	if a>100{
//		return
//	}
//	input:=make([]string,a)
//	for i:=0;i<a;i++{
//		tmp,_:=reader.ReadString('\n')
//		input[i]=strings.TrimSpace(tmp)
//	}
//	res:=0
//	for i:=0;i<len(input);i++{
//		field:=strings.Split(input[i]," ")
//		if len(field)!=2*(i+1)-1{
//			return
//		}
//		tmp:=0
//		for _,v:=range field{
//			number,_:=strconv.Atoi(v)
//			if number>tmp{
//				tmp=number
//			}
//		}
//		res+=tmp
//	}
//	fmt.Println(res)
//}



func main(){
	var a int
	reader:=bufio.NewReader(os.Stdin)
	fmt.Scan(&a)
	if a>100{
		return
	}

	input:=make([]string,a)
	for i:=0;i<a;i++{
		tmp,_:=reader.ReadString('\n')
		input[i]=strings.TrimSpace(tmp)
	}
	arry:=area(a,2*3-1)
	row:=deal1(input,arry)
	fmt.Println(row)
}
func area(m, n int) [][]int {
	a := make([][]int, m)
	for k, _ := range a {
		a[k] = make([]int, n)
	}
	return a
}

func deal1(input []string,arry [][]int)[][]int{
		flag:=0
		for i:=len(input)-1;i>0;i--{
			field:=strings.Split(input[i]," ")
			fmt.Println(field)
			for j:=len(arry)-1;j>0;j--{
				for k:=0;k<len(arry[j]);k++{
					s:=0
					for s<len(field)-1{
						arry[j][s+k],_=strconv.Atoi(field[s])
						s++
						fmt.Println(field[s])
					}
				}
				flag++
			}
		}
		return arry
}