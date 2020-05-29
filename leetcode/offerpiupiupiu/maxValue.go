func maxValue(grid [][]int) int {
	if len(grid)==0||len(grid[0])==0||grid==nil{
		return -1
	}
	m:=len(grid)
	n:=len(grid[0])
	newArry:=make([][]int, m)
	for i:=0;i<m;i++{
		newArry[i]=make([]int,n)
	}


	newArry[0][0]=grid[0][0]

	for i:=1;i<m;i++{
		newArry[i][0]=newArry[i-1][0]+grid[i][0]
	}

	for j:=1;j<n;j++{
		newArry[0][j]=newArry[0][j-1]+grid[0][j]
	}

	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			newArry[i][j]=grid[i][j]+max(newArry[i-1][j],newArry[i][j-1])
		}
	}

return newArry[m-1][n-1]

}  





func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}