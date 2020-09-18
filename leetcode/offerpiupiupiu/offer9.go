package main


//2020717


type CQueue struct {
	in []int
	out []int

}


func Constructor() CQueue {
	queue:=CQueue{}
	return  queue
}


func (this *CQueue) AppendTail(value int)  {
	//传进来的值给out
	this.in=append(this.in,value)
}


func (this *CQueue) DeleteHead() int {
	//如果out的里面没有值，就把in里面的值弹出来给他
	if len(this.out)==0&&len(this.in)==0{
		return -1
	}
	if len(this.out)==0{
		//拿一个值来存这个出栈的值
		for len(this.in)>0{
			tmp:=this.in[len(this.in)-1]
			this.out=append(this.out,tmp)
			this.in=this.in[:len(this.in)-1]
		}
	}
	res:=this.out[len(this.out)-1]
	this.out=this.out[:len(this.out)-1]
	return res
}

