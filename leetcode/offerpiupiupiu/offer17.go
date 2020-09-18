package main


func printNumbers(n int) []int {
	lenth:=1
	res:=[]int{}
	for i:=0;i<n;i++{
		lenth*=10
	}
	for j:=1;j<lenth;j++{
		res=append(res,j)
	}
	return res
}