/*

func findNumberIn2DArray(matrix [][]int, target int) bool {


for j:=0;j<len(matrix[0]);j++{
	if matrix[0][j]>target{
		for i=0;;i++{
			if matrix[i][j-1]==target{
				return true
			}
		}
		return false
	}
}
for k:=0;k<len(matrix);k++{
	if matrix[k][len(matrix[0])]==target{
		return true
	}
}
return false
}

*/

/*

func findNumberIn2DArray(matrix [][]int, target int) bool {
	var right *int
	var up *int
	var now *int
	now=&matrix[0][len(matrix[0])]
	right=&matrix[0][len(matrix[0])+1]
	up=&matrix[1][len(matrix[0])]
	for{
	if *now==target{
		return true
	}
	if *now==nil{
			return false
		}
	if now<target{
		i=*up
	}
	if now>target{
		i=*right
	}
	right=now
}
}

*/

func findNumberIn2DArray(matrix [][]int, target int) bool {
	i:=len(matrix)-1
	j:=0

	for i>-1{
		if j<len(matrix[i]){
		if matrix[i][j]==target{
			return true
		}
		if matrix[i][j]<target{
			j--
		}
		if matrix[i][j]>target{
			i++
		}
	}else{
		return false
	}
		
}
return false
}


func findNumberIn2DArray(matrix [][]int, target int) bool {
    //以左下角为原点
    i:=len(matrix)-1//获取右下角y坐标
    j:=0//获取右下角x坐标
    for i>-1{
        if j<len(matrix[i]){
            if target<matrix[i][j]{
                i--  //小于target,向上查找
            }else if target>matrix[i][j]{
                j++ //大于targat,向右查找
            }else if target==matrix[i][j]{
                return true
            }
        }else{
            return false//超出数组返回false
        }
    }
    return false//超出matrix返回false
}

作者：sakura-151
链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/solution/go-liang-chong-jie-fa-by-sakura-151/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。