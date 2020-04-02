func numWays(n int) int {
	if n==1||n==2{
		return n
	}
	//这里利用了传入的n建立切片，不浪费空间
	s:=make([]int,n+3)
	s[0] = 1 
	s[1] = 1
    s[2] = 2
	for i:=3;i<=n;i++{
		//最小的10位质数，防溢出
		s[i]=(s[i-1]+s[i-2])%1000000007
	}
	return s[n]
}