// 滑动窗口解决连续数
func findContinuousSequence(target int) [][]int {
	res:=make([][]int, 0)
	i,j:=1,1
	//边界不会超过目标值的一半
	broder := (target+1)/2
	for i<=broder&&i<=j{
		sum:=(i+j)*(j-i+1)/2
		//把和与目标值相比
		if sum<target{
			j++
		}else if sum>target{
			i++
		}else{
			temp:=make([]int, 0)
			for k:=i;k<=j;k++{
				temp=append(temp,k)
			}
			res=append(res,temp)
			//加完之后记得要把左边值+1
			i++
		}
	}
	return res
}