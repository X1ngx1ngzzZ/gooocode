package main

import (
	"fmt"
	"reflect"
)

//go支持反射
//把一些数据协议里面的给到go

//reflect demo
func main() {
	a := 1.235
	reflectType(a)
	b := "sdsda"
	reflectType(b)
}

func reflectType(x interface{}) {
	//1、用类型断言,switch
	//2、用反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj)
}
