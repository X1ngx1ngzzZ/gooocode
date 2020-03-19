func climbStairs(n int) int {
if n==1{
	return 1
}	
if n==2{
	return 2
}
t:=climbStairs(n-1)+climbStairs(n-2)
return t
}


func climbStairs(n int) int {
	if n==1{
		return 1
	}	
	if n==2{
		return 2
	}
	first:=1
	second:=2
    var third int
	for i:=3;i<=n;i++{
		third=first+second
		first = second
        second = third 
	}
	return third
}

func climbStairs(n int) int {
	if n==1{
		return n
	}else if n==2{
		return n
	}
	first:=1
	second:=2
	//third没定义
	for i:=3;i<=n;i++{
		third:=first+second
		second=third
		first=second
	}
	return third
}