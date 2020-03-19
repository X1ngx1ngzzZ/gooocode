type CQueue struct {
	A []int
	B []int
}


func Constructor() CQueue {
	Q:=CQueue{}
	Q.A=make([]int,0,5)
	Q.B=make([]int,0,5)
}


func (this *CQueue) AppendTail(value int)  {
    this.A=append(this.A[],value)
}


func (this *CQueue) DeleteHead() int {
    
}
