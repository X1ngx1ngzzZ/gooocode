package main

/*
type MinStack struct {
	data []int
}

func Constructor() MinStack {
	s := MinStack{}
	s.data = make([]int, 50, 100)
	return s
}



func (this *MinStack) Push(x int)  {

	s=append(s,x)
}


func (this *MinStack) Pop()  {
    s=s[:len(s)-1]

func (this *MinStack) Top() int {
	tp:=s[len(s)]
	return tp
}

func (this *MinStack) GetMin() int {
	min:=s[len(s)]
	for k,v:=range s{
		if s[i]<min{
			min=s[i]
		}
	}
	return min
}


//最蠢的方法

func main() {
	fmt.Println(Constructor)
}

*/
/*
type MinStack struct {
	data []int
}

func Constructor() MinStack {
	s := MinStack{}
	s.data = make([]int, 50, 100)
	return s
}



func (this *MinStack) Push(x int)  {
	this.data=append(this.data,x)
}


func (this *MinStack) Pop()  {
    this.data=this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	tp:=this.data[len(this.data)-1]
	return tp
}

func (this *MinStack) GetMin() int {
	min:=this.data[len(this.data)-1]
	for _,v:=range this.data{
		if v<min{
			min=v
		}
	}
	return min
}
*/
/*
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//错误示范
/*
 type MinStack struct {
	data int
	min int
}




/** initialize your data structure here.
func Constructor() MinStack {
	s:=[]MinStack{}
	return s
}

*/

type MinStack struct {
	s []node
}
type node struct {
	data int
	min  int
}

//initialize your data structure here.
func Constructor() MinStack {
	//	s := MinStack{}
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	d := node{data: x, min: x}
	//在这里判断栈是否为空
	if len(this.s) > 0 && this.s[len(this.s)-1].min < x {
		d.min = this.s[len(this.s)-1].min
	}
	this.s = append(this.s, d)

}

func (this *MinStack) Pop() {
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1].data
}

func (this *MinStack) Min() int {
	return this.s[len(this.s)-1].min
}

//试验
/*
func main() {
	a := slice()
	fmt.Println(a[len(a)-1 : len(a)])
}
func slice() []int {
	s := make([]int, 0, 5)
	s = append(s, 5)
	s = append(s, 4)
	return s
}
*/
