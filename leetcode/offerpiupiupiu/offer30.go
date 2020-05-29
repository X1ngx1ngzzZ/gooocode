type MinStack struct {
	data []Node
}


type Node struct{
	val int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	//生成一个节点
	tmp:=Node{val:x,min:x}
	//如果X比前一个节点的最小值大的话，就把自己的Min
	//换成前一个最小值
	if len(this.data)>0&&x>this.data[len(this.data)-1].min{
		tmp.min=this.data[len(this.data)-1].min
	}
	this.data=append(this.data,tmp)
}


func (this *MinStack) Pop()  {
	this.data=this.data[:len(this.data)-1]
}


func (this *MinStack) Top() int {
	return this.data[len(this.data)-1].val
}


func (this *MinStack) Min() int {
	return this.data[len(this.data)-1].min
}
