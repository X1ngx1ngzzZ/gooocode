
//2020718

func pathSum(root *TreeNode, sum int) [][]int {
	if root==nil{
		return nil
	}
	var res [][]int
	dfs(root,sum,[]int{},&res)
	return res

}

func dfs(root *TreeNode,sum int,arry []int,res *[][]int){
	//传进来是空的的话直接返回，
	if root==nil{
		return
	}
	//把传进来根节点的值先加到自己的路径数组里
	arry=append(arry,root.Val)
	//因为每次都是传进来sum-root.Val,所以直接看相等否
	//同时还要看根节点的左右节点都为空，因为路径是到叶子节点
	if root.Val==sum&&root.Left==nil&&root.Right==nil{
		//如果不做拷贝的话，arry指向同一片内存，修改arry直接修改了所有的arry
		tmp:=make([]int,len(arry))
		copy(tmp,arry)
		*res=append(*res,tmp)
}
	dfs(root.Left,sum-root.Val,arry,res)
	dfs(root.Right,sum-root.Val,arry,res)

	//回溯
	arry=arry[:len(arry)-1]
}