type node struct{
	val int
	min int
}



type CQueue struct {
	in 	[]node
	out []node
}





/** initialize your data structure here. */
func Constructor() MinStack {
return 	MinStack{}
}


func (this *MinStack) Push(x int)  {
	newnode:=node{val:x,min:x}
	if len(this.s)!=0{
		if x>this.s[len(this.s)-1].min{
			newnode.min=this.s[len(this.s)-1].min
		}
	}
	this.s=append(this.s,newnode)
}


func (this *MinStack) Pop()  {
	this.s=this.s[:len(this.s)-1]
}


func (this *MinStack) Top() int {
	return this.s[len(this.s)-1].val
}


func (this *MinStack) Min() int {
	return this.s[len(this.s)-1].min
}