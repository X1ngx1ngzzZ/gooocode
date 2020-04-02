//传统方法，非常慢
func fib(n int) int {
if n==0{
	return 0
}else if n==1{
	return 1
}else{
	return fib(n-1)+fib(n-2)
}
}


func fib(n int) int {
	if n==0||n==1{
		return n
	}
	//这里利用了传入的n建立切片，不浪费空间
	s:=make([]int,n+1)
	s[0] = 0 
	s[1] = 1
	for i:=2;i<=n;i++{
		//最小的10位质数，防溢出
		s[i]=(s[i-1]+s[i-2])%1000000007
	}
	return s[n]
}