package main


func lastRemaining(n int, m int) int {
	index:=0
	for i:=2;i<=n;i++{
		index=(index+m)%i
	}
	return index
}