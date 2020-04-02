func movingCount(m int, n int, k int) int {
//先用这个m,n建一个二维数组
arry:=make([][]int, m)
for i,_:=range arry{
	arry[i]=make([]int, n)
}

return deal(m,n,0,0,k,arry)
}

func deal(m,n,i,j,k int,arry [][]int)int{
	if i<0||j<0||i>=m|| j >= n || arry[i][j] == 1 ||(sum(i)+sum(j)) > k{
		return 0
	}
	arry[i][j]=1
	sum:=1
	//这里机器人只往右下走，所以可以根据情况选择
	sum += deal(m, n, i, j+1, k,arry )
	sum += deal(m, n, i, j-1, k, arry)
	sum += deal(m, n, i+1, j, k, arry)
	sum += deal(m, n, i-1, j, k, arry)
	return sum
}







func sum(n int)int{
	res:=0
	for n>0{
		res+=n%10
		n=n/10
	}
	return res
}