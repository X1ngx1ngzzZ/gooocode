package main

import (
	"fmt"
	"log"
)

type Stack struct {
	size int
	top  int
	data []interface{}
}

//判断长度
func (t *Stack) lens() int {
	return t.top
}

//清空
/*
func (t *Stack) clear() {
	t.top == 0
}
*/

//判断是否为空
func (t *Stack) IsEmpty() bool {
	return t.top == 0
}

//判满
func (t *Stack) IsFull() bool {
	return t.top == t.size
}

//遍历
func (t *Stack) Traverse(fn func(node interface{}), zheng bool) {
	if zheng {
		for i := 0; i < t.top; i++ {
			fn(t.data[i])
		}
	} else {
		for i := t.top - 1; i >= 0; i-- {
			fn(t.data[i])
		}
	}
}

//初始化
func InitStack(size int) Stack {
	q := Stack{}
	q.size = size
	q.data = make([]interface{}, size)
	return q
}

//入栈
func (t *Stack) Push(element interface{}) bool {
	if t.IsFull() {
		fmt.Println("满了")
		return false
	}
	t.data[t.top] = element
	t.top++
	return true
}

//出栈
func (t *Stack) pop() (r interface{}, err error) {
	if t.IsEmpty() {
		err = fmt.Errorf("栈已满，无法完成入栈")
		log.Println("栈已满，无法完成入栈")
		return
	}
	t.top--
	r = t.data[t.top]
	return
}

func isValid(s string) bool {
	stack := InitStack(64)
	m := make(map[rune]int)
	m = map[rune]int{
		'(': 1,
		')': 2,
		'[': 3,
		']': 4,
		'{': 5,
		'}': 6,
	}

	if s == "" {
		return true
	} else {
		for _, v := range s {
			if m[v] == 1 || m[v] == 3 || m[v] == 5 {
				stack.Push(v)
			} else {
				switch m[v] {
				case 2:
					if stack.data[stack.top-1] == '(' {
						stack.pop()
					} else {
						return false
					}
				case 4:
					if stack.data[stack.top-1] == '[' {
						stack.pop()
					} else {
						return false
					}
				case 6:
					if stack.data[stack.top-1] == '{' {
						stack.pop()
					} else {
						return false
					}
				}
			}
		}
		if stack.IsEmpty() {
			return true
		} else {
			return false
		}
	}

}

func main() {
	s := "()}"
	fmt.Println(isValid(s))
}


func isValid(s string) bool {
	m:=make(map[rune](rune)){'(':')','[':']','{','}'}

	stack:=make[]rune

	for k,v:=range s{
		if v=='('&&v=='['&&v=='{'{
			stack=append(stack,v)
		}else if len(stack)>0&&m[v]==Stack[:len(stack)-1]{
			stack=stack[:len(stack)-1]
		}else{
			return false
		}

	}
	return len(stack)==0
}


func isValid(s string) bool {
    brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
    var stack []rune

    for _, char := range s {
        fmt.Println(reflect.TypeOf(char))
        if char == '(' || char == '{' || char == '[' {
            // 入栈
            stack = append(stack, char)
            // 循环中，stack不能为空
        } else if len(stack) > 0 && brackets[char] == stack[len(stack) - 1] {
            // 栈中有数据，且此元素与栈尾元素相同
            stack = stack[:len(stack) - 1]
        } else {
            return false
        }
    }

    // 循环结束，栈中还有数据则 false
    return len(stack) == 0
}


func isValid(s string) bool {
	//m:=map[rune]rune{'(':')','[':']','{':'}'}
    m:=map[rune]rune{')':'(',']':'[','}':'{'}

	stack:=make([]rune,0,5)

	for _,v:=range s{
		if v=='('||v=='['||v=='{'{
			stack=append(stack,v)
		}else if len(stack)>0&&m[v]==stack[len(stack)-1]{
			stack=stack[:len(stack)-1]
		}else{
			return false
		}

	}
	return len(stack)==0
}
