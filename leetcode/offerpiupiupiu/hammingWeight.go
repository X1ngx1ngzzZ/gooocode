func hammingWeight(num uint32) int {
	res:=0
	num1:=fmt.Sprintf("%b",num)
	for _,v:=range num1{
		if v == '1'{
		res++
	}
}
	return res
}

func hammingWeight(num uint32) int {
    n:=0
    nums:=fmt.Sprintf("%b",num)
    for _,char:=range nums{
        if char=='1'{
            n++
        }
    }
    return n
}

