//二叉树相关 1为递归相关，2为非递归
//遍历共分，前中后序和宽度优先
package main

import "fmt"

type tree struct{
	Val int
	left *tree
	right *tree
}

//创建二叉树！不会不会不会






//前序遍历递归实现，前中后序遍历的代码只有顺序不同，指的是根节点输出的时间点
func qianxu1(Tree *tree){
	fmt.Println(Tree.Val)
	qianxu1(Tree.left)
	qianxu1(Tree.right)

}


//中序非递归
func zhongxu2(Tree *tree) []int{
	if Tree == nil{
		return nil
	}
	result:=[]int
	stack:=[]*tree
	for Tree != nil|| len(stack) != 0{
		if Tree !=nil{
			stack = append(stack,Tree)
			Tree = Tree.left
		}else{
			Tree = stack[:len(stack)-1]
			result = append(result,Tree.Val)
			stack = stack[:len(stack)-1]
			Tree = Tree.right

		}
	}
	return result

}

//前序非递归
func qianxu2(Tree *tree) []int{
	if Tree == nil{
		return nil
	}
	result:=[]int
	stack:=[]*tree
	for Tree != nil|| len(stack) != 0{
		if Tree !=nil{
			//压栈的时候就直接加到结果里
			result = append(result,Tree.Val)
			stack = append(stack,Tree)
			Tree = Tree.left
		}else{
			Tree = stack[:len(stack)-1]
			stack = stack[:len(stack)-1]
			Tree = Tree.right

		}
	}
	return result

}

func main(){
	start:=[]int{1,2,3,4,5,6,7,8,9,10}
	tree:=TreeCreate(3,start)
}