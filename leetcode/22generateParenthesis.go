func generateParenthesis(n int) []string {
	res := make([]string, 0)
	deal(&res,"",n,n)
	return res
}

func deal(res *[]string,tmp string,left ,right int){
	if left==0{
		for i:=0;i<right;i++{
			tmp+=")"
		}
	*res=append(*res,tmp)
	return
	}
	deal(res,tmp+"(",left-1,right)
	if left<right{
		deal(res,tmp+")",left,right-1)
	}
}