type node struct{
	val int
	max int
}

type CQueue struct {
	in 	[]int
	out []int
}


type MaxQueue struct {
    in 	[]node
	out []node
}


func Constructor() MaxQueue {
    return Max_value{}
}


func (this *MaxQueue) Max_value() int {
	if len(this.out)==0{
		for len(this.in)>0{
			tmp:=this.in[len(this.in)-1]
			this.out=append(this.out,tmp)
			this.in=this.in[:len(this.in)-1]
		}
    }
	return this.out[len(this.s)-1].max
}


func (this *MaxQueue) Push_back(value int)  {
	tmp:=node{val:value,max:value}
	if len(this.in)!=0{
		if value>this.in[len(this.s)-1].max{
			tmp.max=this.in[len(this.s)-1].max
		}
	}
	this.in=append(this.in,tmp)
}


func (this *MaxQueue) Pop_front() int {
    var res int
	if len(this.out)==0{
		for len(this.in)>0{
			tmp:=this.in[len(this.in)-1]
			this.out=append(this.out,tmp)
			this.in=this.in[:len(this.in)-1]
		}
    }
		res=this.out[len(this.out)-1].val
		this.out=this.out[:len(this.out)-1]
		return res

}


/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */