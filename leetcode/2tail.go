type CQueue struct {
	in 	[]int
	out []int
}


func Constructor() CQueue {
	return CQueue{}
}


func (this *CQueue) AppendTail(value int){
	//访问结构体里的记得加.
	this.in = append(this.in,value)
}


func (this *CQueue) DeleteHead() int {
	//特殊情况
	if len(this.in)==0&&len(this.out)==0{
		return -1
	}
	//out栈为空，把in的遍历放进去
	if len(this.out)==0{
		for len(this.in)>0{
			invalue:=this.in[len(this.in)-1]
			this.in = this.in[：len(this.in)-1]
			this.out = append(this.out,invalue)
		}
	}

	//出out栈顶的元素
	outvalue:=this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return outvalue
}