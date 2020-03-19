func isPalindrome(x int) bool {
	y:=0
	if x<0{
	 return false
 }else{
	 //拿数保存他x
	 a:=x
	for x!=0{
		y=y*10+x%10
		x/=10
	}
	retrun y==a
 }
}







func isPalindrome(x int) bool {
	y:=0
	for x!=0{
		y=y*10+x%10
		x=x/10
	}
    if y!=x{
        return false
    }else{
        return true
    }
}
//反转一半的数字
func isPalindrome(x int) bool {
    if x<0||x%10==0&&x!=0{
        return false
    }

    revertedNumber := 0
    for x > revertedNumber {
            revertedNumber = revertedNumber * 10 + x % 10
            x /= 10
        }
//当x<=revertedNumber, 则反转已达数字长度一半
        return x == revertedNumber || x == revertedNumber/10
}