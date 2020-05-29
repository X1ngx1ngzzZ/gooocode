//构造结构体
package main

import "fmt"

//Person 结构体
type Person struct {
	name string
	city string
	age  int
}

func newPerson(name, city string, age int) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	p1 := newPerson("fjj", "nanning", 20)
	fmt.Println(p1)
	p1.setnewage(18)
	fmt.Println(p1)
}

func (p *Person) setnewage(newage int) {
	p.age = newage
}

/*
什么时候应该使用指针类型接收者
需要修改接收者中的值
接收者是拷贝代价比较大的大对象
保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/
