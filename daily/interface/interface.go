package main

import "fmt"

func main() {
	var x interface{}

	x = "heel"
	fmt.Println(x)
	x = 3
	fmt.Println(x)
	x = true
	fmt.Println(x)
	//空接口 类型断言
	ret, ok := x.(bool)
	if !ok {
		fmt.Println("Not bool")
	} else {
		fmt.Println(ret, "is bool")
	}
	//可以使用switch语句进行类型断言 

}

//空接口类型可以做map的value
