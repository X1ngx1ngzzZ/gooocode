package main


//2020716
func maxSlidingWindow(nums []int, k int) []int {
	if nums==nil||k==0||len(nums)<k{
		return nil
	}
	var tmp int
	var res []int
	i:=0
	j:=k
	for j<=len(nums){
		tmp=nums[i]
		for k:=i;k<j;k++{
			if nums[k]>tmp{
				tmp=nums[k]
			}
		}
		res=append(res,tmp)
		i++
		j++
	}
	return res
}

//暴力法优化
func maxSlidingWindow(nums []int, k int) []int {
	var max int
	var res []int

	for i,j:=0,k-1;j>=0&&j<len(nums);j++{
		//如果max刚出去，或者是第一个的话
		if i==0||max==nums[i-1]{
			max=nums[i]
			//for z:=i;z<j;z++ {
			//	if max<nums[z] {
			//		max = nums[z]
			//	}
			//}
			for t := j; t > i; t-- {
				if max < nums[t] {
					max = nums[t]
				}
			}
		}else{
			if nums[j]>max{
				max=nums[j]
			}
		}
		res=append(res,max)
		i++
	}
	return res
}

//II
type MaxQueue struct {
	que []int
	max []int
}


func Constructor() MaxQueue {
	return MaxQueue{
		que: make([]int,0),
		max: make([]int,0),
	}
}


func (this *MaxQueue) Max_value() int {
	if len(this.max)==0{
		return  -1
	}
	return this.max[0]
}


func (this *MaxQueue) Push_back(value int)  {
	this.que=append(this.que,value)
	for len(this.max)!=0&&value>this.max[len(this.max)-1]{
		this.max=this.max[:len(this.max)-1]
	}
	this.max=append(this.max,value)

}


func (this *MaxQueue) Pop_front() int {
	n:=-1
	if len(this.que)!=0{
		n=this.que[0]
		this.que=this.que[1:]
		if n==this.max[0]{
			this.max=this.max[1:]
		}
	}
	return n
}


