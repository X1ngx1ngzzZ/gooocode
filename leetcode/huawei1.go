package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)


func main(){
	reader:=bufio.NewReader(os.Stdin)
	input,_:=reader.ReadString('\n')
	fieldRaw:=strings.Fields(input)
	field:=make([]int,2)
	for i:=0;i<2;i++{
		//tmp,_:=strconv.Atoi(fieldRaw[i])                0 1 2
		field[i],_=strconv.Atoi(fieldRaw[i])
		if field[i]<10||field[i]>1000{
			kong:=make([][]int,0)
			fmt.Println(kong)

		}
	}

	//建立了矩阵
	team:=area(field[0],field[1])
	//fmt.Println(team)
	//给矩阵标值
	rawarry:=giveValue(team)
	//fmt.Println(rawarry)
	//顺时针打印值得到一个数组
	res:=spiralOrder(rawarry)
	//fmt.Println(res)
	//判断这个数组里面的值是不是符合要求，符合再留下
	fmt.Println(deallast(res))

}
func area(m, n int) [][]int {
	a := make([][]int, m)
	for k, _ := range a {
		a[k] = make([]int, n)
	}
	return a
}
func areastring(m, n int) [][]string {
	a := make([][]string, m)
	for k, _ := range a {
		a[k] = make([]string, n)
	}
	return a
}
func giveValue(arry [][]int)[][]string{
	arrystring:=areastring(len(arry),len(arry[0]))
	for i:=0;i<len(arry);i++{
		for j:=0;j<len(arry[i]);j++{
			tmp:=fmt.Sprintf("%d,%d",i,j)
			arrystring[i][j]=tmp
		}
	}
	return  arrystring
}
func spiralOrder(matrix [][]string) []string {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	const up, down, left, right = 0, 1, 2, 3
	i, j := len(matrix), len(matrix[0])
	s, x, z, y := 0, i-1, 0, j-1
	r, c, direct := 0, 0, right
	result := make([]string, i*j)

	for k := 0; k < len(result); k++ {
		//tmp:=[]int{r,c}
		result[k] = matrix[r][c]
		//result[k]=tmp
		switch direct {
		case right:
			if c < y {
				c++
			} else {
				direct = down
				s++
				r++
			}
		case down:
			if r < x {
				r++
			} else {
				direct = left
				y--
				c--
			}
		case left:
			if c > z {
				c--
			} else {
				direct = up
				x--
				r--
			}
		case up:
			if r > s {
				r--
			} else {
				direct = right
				z++
				c++
			}
		}
	}

	return result
}
func deallast(arry []string)[][]string{
	res:=[][]string{}
	for k:=0;k<len(arry);k++{
		if k%10==7&&k/10%10%2 == 1{
				//fmt.Println(k)
				//fmt.Println(arry[k])
				tmp := strings.Split(arry[k-1], ",")
				middle := make([]int, 2)
				for i := 0; i < 2; i++ {
					middle[i], _ = strconv.Atoi(tmp[i])
				}
				last := []string{arry[k-1]}
				res = append(res, last)
			}
		}

	return res
}