func maxProfit(prices []int) int {
	temp:=0
	profile:=0
	n:=len(prices)-1
	for i:=1;i<=n;i++{
		temp=prices[i]-prices[i-1]
		if temp>0{
			profile+=temp
		}
	}
	return profile
}