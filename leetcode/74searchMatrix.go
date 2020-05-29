//比每一行最后一个值确定行数
//依次找这一行
//或者二分法也行
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0||len(matrix[0])==0{
		return false
	}
	hang:=len(matrix)-1
	lie:=len(matrix[0])-1
	for i:=0;i<=hang;i++{
		if matrix[i][lie]>=target{
			for j:=0;j<=lie;j++{
				if matrix[i][j]==target{
					return true
				}
			}
			return false
		}
	}
	return false
}