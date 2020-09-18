package main

import "fmt"

func main(){
	deal(2178)
}
func deal(i int){
	fan:=reverse(i)
	fmt.Println(fan)
	tmp:=i*4
	fmt.Println(tmp)

	if tmp==fan{
		println("yes")
	}
}

func reverse(x int) int {
	y := 0
	for x!=0 {
		y = y*10 + x%10
		if !( -(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}