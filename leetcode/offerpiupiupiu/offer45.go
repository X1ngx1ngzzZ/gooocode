func minNumber(nums []int) string {
	if len(nums)==0{
		return ""
	}
	out:=""
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if bigger(nums[i],nums[j])==nums[i]{
				nums[i],nums[j]=nums[j],nums[i]
			}
		}
	}

	for i:=0;i<len(nums);i++{
		out+=strconv.Itoa(nums[i])
	}

	return out
	
}

func bigger(a,b int)int{
a1:=strconv.Itoa(a)
a2:=strconv.Itoa(b)
z:=a1+a2
f:=a2+a1
if z>f{
	return a
}else{
	return b
}
}



//2020718
func minNumber(nums []int) string {
	if len(nums)==0{
		return ""
}
	res:=""
	//冒泡
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			//change
			if bigger(nums[i],nums[j]){
				nums[i],nums[j]=nums[j],nums[i]
}
}
}

	for i:=0;i<len(nums);i++{
	res+=strconv.Itoa(nums[i])
	}
	return res
}


func bigger(a,b int)bool{
	a1:=strconv.Itoa(a)
	a2:=strconv.Itoa(b)
	pre:=a1+a2
	last:=a2+a1
	if pre>last{
		return true
}else{
	return false
}
}