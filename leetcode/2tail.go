package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CQueue struct {
	in  []int
	out []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	//访问结构体里的记得加.
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() {
	//特殊情况
	/*
		if len(this.in) == 0 && len(this.out) == 0 {
			return
	*/
	//out栈为空，把in的遍历放进去
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			invalue := this.in[len(this.in)-1]
			this.in = this.in[:len(this.in)-1]
			this.out = append(this.out, invalue)
		}
	}

	//出out栈顶的元素
	//	outvalue := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	//	return outvalue
}

func (this *CQueue) search() int {
	if len(this.in) == 0 && len(this.out) == 0 {
		return -1
	}
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			invalue := this.in[len(this.in)-1]
			this.in = this.in[:len(this.in)-1]
			this.out = append(this.out, invalue)
		}
	}
	return this.out[len(this.out)-1]
}

func main() {
	var T int
	var s1 []string
	var Q1 CQueue
	fmt.Println("input your operations")
	fmt.Scanf("%d", &T)
	fmt.Printf("%d", T)
	s := make([]string, T, 100)
	for i := 0; i < T; i++ {
		fmt.Scanf("%s", &s[i])

	}
	for _, v := range s {
		s1 = strings.SplitN(v, " ", 2)
		if s1[1] != " " {
			value, _ := strconv.Atoi(s1[1])
			Q1.AppendTail(value)
			fmt.Printf("%d is been add", value)
		} else {
			switch s1[0] {
			case "pool":
				Q1.DeleteHead()
				fmt.Println("deleted")
			case "peek":
				result := Q1.search()
				fmt.Printf("%d", result)
			}
		}
	}

}
