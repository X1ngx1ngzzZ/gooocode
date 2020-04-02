func minArray(numbers []int) int {
	return deal(numbers,0,len(numbers)-1)
}


func deal(numbers []int,low int,high int)int{
	//有序的话，旋转为0也算
	if numbers[low]<numbers[high]{
		return numbers[low]
	}
	var m int
	//这个大于等于和low<high很关键
	for numbers[low]>=numbers[high]&&low<high{
		m=(high+low)/2
		if numbers[m]>numbers[low]{
			low=m+1
		}else if numbers[m]<numbers[low]{
			high=m
		}else{
			low++
		}
	}
	return numbers[low]
}




func minArray(numbers []int) int {
    //Left为数组最左端下标，Right为数组最右端下标
    Left:=0
    Right:=len(numbers)-1
    for Left<Right{
        mid:=(Left+Right)/2
        //有序数组
        if numbers[Left]<numbers[Right]{
            return numbers[Left]
        }
        //如果numbers[mid]大于numbers[Left],说明mid左边是递增有序数组
        //而又因numbers[Left]大于numbers[Right]，说明最小值不在左边
        //所以舍弃包括mid左边的子数组
        if numbers[mid]>numbers[Left]{
            Left=mid+1
        }else if numbers[mid]<numbers[Left]{
            Right=mid
        //如果numbers[mid]==numbers[Left]，则移动左边的下标即可
        }else{
            Left++
        }
    }
    return numbers[Left]
}


 