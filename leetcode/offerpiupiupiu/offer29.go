func spiralOrder(matrix [][]int) []int {
	if len(matrix)==0||len(matrix[0])==0{
		return nil
	}
	const up,down,left,right =0,1,2,3
	i,j:=len(matrix),len(matrix[0])
	s,x,z,y:=0,i-1,0,j-1
	r,c,direct:=0,0,right
	result:=make([]int,i*j)
	for k:=0;k<len(result);k++{
		result[k]=matrix[r][c]
		switch direct {
		case right:
			if c<y{
				c++
			}else{
				direct=down
				s++
				r++
			}
		case down:
			if r<x{
				r++
			}else{
				direct=left
				y--
				c--
			}
		case left:
			if c>z{
				c--
			}else{
				direct=up
				x--
				r--
			}
		case up:
			if r>s{
				r--
			}else{
				direct=right
				z++
				c++
			}
		}
	}

	return result

}





func spiralOrder(matrix [][]int) []int {
	if len(matrix)==0||len(matrix[0])==0{
		return nil
	}
	i,j:=len(matrix),len(matrix[0])
	//判断往哪走
	const up,down,left,right=1,2,3,4 
	//限制走到哪一圈
	s,x,z,y:=0,i-1,0,j-1
	result:=make([]int,i*j)
	r,c,direct:=0,0,right
	for k:=0;k<len(result);k++{
		result[k]=matrix[r][c]
		switch direct{
		case right:
			//这里由于是先j++的，把最后一列赋值给过了，所以不需要+等于号
			//同理由于最后一列最后一个记录过了，所以j--
			if c<y{
				c++
			}else{
				direct=down
				s++
				r++
			}
		case down:
			if r<x{
				r++
			}else{
				direct=left
				y--
				c--
			}
		case left:
			if c>z{
				c--
			}else{
				direct=up
				x--
				r--
			}
		case up:
			if r>s{
				r--
			}else{
				direct=right
				z++
				c++
			}
		}
		
	
	}
    return result
}