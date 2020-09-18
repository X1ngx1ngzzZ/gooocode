package main


import (
	"bufio"
	"fmt"
	"math"
	"os"
)


func main(){
	reader:=bufio.NewReader(os.Stdin)
	input:=make([]string,2)
	for i:=0;i<2;i++ {
		input[i],_=reader.ReadString('\n')
	}
	result:=deal(input[0],input[1])
	fmt.Println(result)
}

func deal(str1,str2 string)int{
	tmp:=0
	min:=len(str2)
	max:=0
	for i:=0;i<len(str2);i++{
		for j:=0;j<len(str1);j++{
			if str2[i]==str1[j]{
				//fmt.Println("j:",j)
				tmp=j
				if tmp>=max{
					max=tmp
					//fmt.Println("max",max)
				}else if tmp<=min{
					min=tmp
					//fmt.Println("min",min)
				}
			}
		}
	}
	res:=max-min+1
	return res
}



func minWindow(s string, t string)int {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if (r - l + 1 < len) {
				len = r - l + 1
				ansL, ansR = l, l + len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}

	//return s[ansL:ansR]
	return ansR-ansL+1
}

func main(){
	reader:=bufio.NewReader(os.Stdin)
	input:=make([]string,2)
	for i:=0;i<2;i++ {
		input[i],_=reader.ReadString('\n')
	}
	result:=minWindow(input[0],input[1])
	fmt.Println(result)
}

func minWindow(s string, t string)int {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if (r - l + 1 < len) {
				len = r - l + 1
				ansL, ansR = l, l + len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return 0
	}

	//return s[ansL:ansR]
	return ansR-ansL+1
}



