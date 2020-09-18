package main

import "container/list"

type LRUCache struct {
	keyMap map[int]*list.Element
	elemlist	*list.List
	capcity		int
}

type Node struct {
	K	 int
	V	 int
}

func Constructor(capacity int) LRUCache {
	var ll list.List
	ll.Init()
	return LRUCache{
		keyMap: make(map[int]*list.Element,capacity),
		elemlist: &ll,
		capcity: capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	if v,ok:=this.keyMap[key];ok{
		this.elemlist.MoveToFront(v)
		res:=v.Value.(Node)
		return res.V
	}else{
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	NewNode:=Node{key,value}
	//不在Map里的时候
	if v,ok:=this.keyMap[key];!ok{
		if this.elemlist.Len()==this.capcity{
			ele:=this.elemlist.Back()
			this.elemlist.Remove(ele)
			//删除map里面的值
			del:=ele.Value.(Node).K
			delete(this.keyMap,del)
		}
		element:=this.elemlist.PushFront(NewNode)
		this.keyMap[NewNode.K]=element
	}else{
		v.Value=NewNode
		this.elemlist.MoveToFront(v)
	}
}
