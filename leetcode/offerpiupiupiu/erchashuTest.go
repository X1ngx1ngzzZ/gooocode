package main

import "fmt"

type treenode struct {
	data  int
	left  *treenode
	right *treenode
}

//增加左节点
func (node *treenode) addLeft(value int) {
	children := treenode{data: value}
	node.left = &children
}

func (node *treenode) addRight(value int) {
	children := treenode{data: value}
	node.right = &children
}

func main() {
	node := treenode{data: 3, left: &treenode{data: 3, left: nil, right: nil}, right: &treenode{data: 5, left: nil, right: nil}}
	node.addLeft(8)
	node.addRight(9)
	traversal(&node)

}
func traversal(node *treenode) {
	if node == nil {
		return
	}
	fmt.Println("数据", node.data)
	traversal(node.left)
	traversal(node.right)
	//fmt.Println("数据", node.data)
}

//非递归的遍历二叉树
func inOrderBinaryTree1(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	stack := []*BinaryTreeNode{}
	top := -1
	for top >= 0 || root != nil {
		//找到最左根节点
		for root != nil {
			//把这条线上的入栈
			stack = append(stack, root)
			//栈里有多少元素
			top++
			root = root.lChild
		}
		//把栈顶的元素给item
		item := stack[top]
		stack = stack[:top]
		top-- // 出栈
		fmt.Print(item.data)
		//判断栈顶的元素有没有右边，如果有右边就把root给他
		if item.rChild != nil {
			root = item.rChild
		}
	}
}
