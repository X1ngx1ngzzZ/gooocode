//剪绳子，可以数学解一下，res=(n/m)^m,求极值点为2.7，得出接近3
//要注意给的n，可能会发生越界情况
func cuttingRope(n int) int {
	if n<=2{
		return 1
	}
	if n==3{
		return 2
	}
	k:=n/3
	v:=n%3
	if v==0{
		result:=math.Pow(3,float64(k))
		
		return  
	}else if v==1{
		result:=math.Pow(3,float64(k-1)) 
		result*=4
		return int(result%1000000007)
	}else{
		result:=math.Pow(3,float64(k))
		result*=2
		return int(result%1000000007)
	}
}