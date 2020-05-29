
 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}
   




//返回一维数组的
func levelOrder(root *TreeNode) []int {
	if root==nil{
		return nil
	}
	res:=make([]int,0)
	//创建个队列，直接把根节点放里面
	queue:=[]*TreeNode{root}
	//队列不为空时
	for len(queue)!=0{
	//这里新建一个队列来模拟出队
	temp:=[]*TreeNode{}	
	for _,v:=range queue{
	res=append(res,v.Val)
		if v.Left!=nil{
			temp=append(temp,v.Left)
		}
		if v.Right!=nil{
			temp=append(temp,v.Right)
		}
		}
	queue=temp
	}
	return res
	
}


//返回二维数组的
func levelOrder(root *TreeNode) [][]int {
	if root==nil{
		return nil
	}
	//创建二维数组存一下
	res:=make([][]int,0)
	//创建个队列，直接把根节点放里面
	queue:=[]*TreeNode{root}
	//队列不为空时
	for len(queue)!=0{
	//这里新建一个队列来模拟出队
	temp:=[]*TreeNode{}
	//建一个一维数组存这一层的
	temparry:=[]int{}	
	for _,v:=range queue{
		temparry=append(temparry,v.Val)
		if v.Left!=nil{
			temp=append(temp,v.Left)
		}
		if v.Right!=nil{
			temp=append(temp,v.Right)
		}
		}
	//遍历完这一层把这个数组存进去
	res=append(res,temparry)
	queue=temp
	}
	return res	
}



//返回二维数组的，之字型打印
func levelOrder(root *TreeNode) [][]int {
	if root==nil{
		return nil
	}
	//创建二维数组存一下
	res:=make([][]int,0)
	//创建个队列，直接把根节点放里面
	queue:=[]*TreeNode{root}
	//建一个flag
	flag:=1
	//队列不为空时
	for len(queue)!=0{
	//这里新建一个队列来模拟出队
	temp:=[]*TreeNode{}
	//建一个一维数组存这一层的
	temparry:=[]int{}	
	for _,v:=range queue{
		if flag%2==1{
		temparry=append(temparry,v.Val)
		}else{
			now:=[]int{v.Val}
		//	temparry=append(now,temparry)这里加上三个.就可以过
		//这种append写法可以把后面这个切片加到前面的前面后面
		//这里就表现为temparry放在now后面，再返回给temparry
		//即完成了头插法
			temparry=append(now,temparry...)
		}
		if v.Left!=nil{
			temp=append(temp,v.Left)
		}
		if v.Right!=nil{
			temp=append(temp,v.Right)
		}
		}
	//遍历完这一层把这个数组存进去
	res=append(res,temparry)
	queue=temp
	flag++
	}
	return res	
}
