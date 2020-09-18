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


//2020617

type Node struct{
	data int
	min int
}

type MinStack struct{
	slice []Node
}

func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
newnode:=Node{data:x,min:x}
if len(this.slice)>0&&this.slice[len(this.slice)-1].min<x{
newnode.min=this.slice[len(this.slice)-1].min
}
this.slice=append(this.slice,newnode)

}


func (this *MinStack) Pop()  {
this.slice=this.slice[:len(this.slice)-1]
}


func (this *MinStack) Top() int {
return this.slice[len(this.slice)-1].data
}


func (this *MinStack) Min() int {
return this.slice[len(this.slice)-1].min
}