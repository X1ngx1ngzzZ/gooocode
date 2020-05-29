//暴力解

abcdcdefg
//slide window
func lengthOfLongestSubstring(s string) int {
	m:=make(map[byte]int)
	res:=0
	for i,j:=0,0;i<len(s);i++{
		//把值存进去
		m[s[i]]++
		for m[s[i]]>1{
			//有值相等
			m[s[j]]--
			j++
		}
		res=max(res,i-j+1)
	}
	return res
}

func max (x, y int) int {
    if x > y {
        return x
    }
    return y
}


//slice
func lengthOfLongestSubstring(s string) int {
	var Length int
	var s1 string
	left := 0
	right := 0
	s1 = s[left:right]
	for ; right < len(s); right++ {

		if index := strings.IndexByte(s1, s[right]); index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > Length {
			Length = len(s1)
		}
	}

	return Length
}

