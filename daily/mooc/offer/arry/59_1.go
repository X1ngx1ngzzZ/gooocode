//暴力法
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


func maxSlidingWindow(nums []int, k int) []int {
	ans,que,k:=[]int{},[]int{},k-1
	for i,v:=range nums{
		for len(que)>0&&que[len(que)-1]<nums{
			que=que[:len(que)-1]
		}
		que =append(que,nums)
		if i>=k{
			ans =append(ans,que[0])
			if que[0]==nums[i-k]{
				que=que[1:]
			}
		}
	}
}