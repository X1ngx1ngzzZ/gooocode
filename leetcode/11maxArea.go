func maxArea(height []int) int {
	i:=0
	j:=len(height)-1
	res:=0
	for i<j{
		
		tmp:= min(height[i],height[j])*(j-i)
		if tmp>res{
			res=tmp
		}
		if height[i]==min(height[i],height[j]){
			i++
		}else{
			j--
		}
	}
	return res
}


func min(a,b int)int{
	if a>b{
		return b
	}else{
		return a
	}
}