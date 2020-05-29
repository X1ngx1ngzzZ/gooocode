
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
		Left:=maxDepth(root.Left)
		Right:=maxDepth(root.Right)
	return max(Left, Right) + 1
	}
	
	func max(a,b int)int{
		if a > b {
			return a
		}
		return b
	}





func maxDepth(root *TreeNode) int {
		if root == nil {
			return 0
		} 
		result := 0
		
		queue := make([]*TreeNode,0)
		queue = append(queue,root)
		for len(queue) > 0{
			level_size := len(queue)
			//把每一层取完再重新获取队列长度
			for level_size > 0 {
				level_size--
				node := queue[0]
				queue = queue[1:]
	
				if node.Left != nil{
					queue = append(queue,node.Left)
				}
				 if node.Right != nil{
					queue = append(queue,node.Right)
				}
			}
			result++
		}
		return result
	}
	
	
func maxDepth(root *TreeNode)int{
	if root ==nil{
		return 0
	}
	queue:=[]*TreeNode{root}
	result:=0
	
	for len(queue)!=0{
		//新队列模仿出队
		temp:=[]*TreeNode{}
		//遍历每一层，把下层添加给新队
		for _,v:=range queue{
			if v.Left!=nil{
				temp = append(temp,v.Left)
			}
			if v.Right!=nil{
				temp = append(temp,v.Right)
			}
		}

	queue=temp
	result++
	}
	return result
}