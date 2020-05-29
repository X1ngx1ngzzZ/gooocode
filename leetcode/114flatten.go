//左不为空，向右进一个
//左为空，左给右，右给左的最右
//左置为空

func flatten(root *TreeNode)  {
	head:=root
	var tmp *TreeNode
	for head!=nil{
		if head.Left==nil{
			head=head.Right
		}else{
			tmp=head.Left
			for tmp.Right!=nil{
				tmp=tmp.Right
			}
			head.Right,tmp.Right=head.Left,head.Right
		head.Left=nil
		}
	}

}