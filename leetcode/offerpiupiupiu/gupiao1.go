func maxProfit(prices []int) int {
	minprices := 1<<31
	maxvalue := 0
	for _,v:=range prices{
		minprices = min(v,minprices)
		maxvalue = max(maxvalue,(v-minprices))
	}
	return maxvalue
}

func min(a,b int)int{
	if a<b{
		return a
	}else{
		return b
	}
}


func max(a,b int)int{
	if a<b{
		return b
	}else{
		return a
	}
}