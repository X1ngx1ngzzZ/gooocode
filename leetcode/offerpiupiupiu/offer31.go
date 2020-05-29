
//这想法很辣鸡实现还麻烦
func validateStackSequences(pushed []int, popped []int) bool {
//	var k,v int
	var i,j int
	slice:=make([]int, 0)
	for k,v=range popped{
		//检查是不是空的
		if len(slice)!=0{
			//不为空就检查栈顶元素，一样就出栈，出栈第一个元素
			if v==slice[len(slice)-1]{
				slice=slice[:len(slice)-1]
			}
			//栈顶元素不同，就遍历A找到这个数字，把它之前的全部入栈
			//如何处理后面的数字
			for i,j=range pushed{
				if j==v{
					break
				}
				slice=append(slice,j)
			}
		}

	}
}








func validateStackSequences(pushed []int, popped []int) bool {
	slice:=make([]int,0)
	//遍历输入数组，依次入栈
	//如果栈栈不为空且栈顶元素和给定的顺序
	//里出栈的元素相等的时候就弹出,循环读取B的元素
	i:=0
	for _,v:=range pushed{
		slice=append(slice,v)
		for len(slice)!=0&&slice[len(slice)-1]==popped[i]{
			slice=slice[:len(slice)-1]
			i++
		}
	}
	if len(slice)==0{
		return true
	}else {
		return false
	}
}