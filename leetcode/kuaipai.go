package main

import "fmt"

func main(){
	slice:=[]int{7,5,3,2,1,3}
	QuickSort(slice,0,len(slice)-1)
	fmt.Println(slice)
}

func QuickSort(arr []int,l,r int){
	if l<r{
		flag:=arr[r]
		i:=l-1
		for j:=l;j<r;j++{
			if arr[j]<=flag{
				i++
				arr[j],arr[i]=arr[i],arr[j]
			}
		}
		i++
		arr[r],arr[i]=arr[i],arr[r]
		QuickSort(arr,l,i-1)
		QuickSort(arr,i+1,r)
	}




}