func shortest(strs []string)string{
	if len(strs)==0{
		return ""
	}
	short:=strs[0]
	for _, v := range strs {
		
		if len(v) <len(short) {
			if len(v)==0{
				return ""
			}
			short= v
		} 
		}
		return short
	}


func longestCommonPrefix(strs []string) string {
	short:=shortest(strs)
	
	//对最短的这个字符串进行遍历
	for i, v := range short {
		for j := 0; j < len(strs); j++ {
			//切片底层是由byte，而V是rune格式，所以强制转换
			if strs[j][i] != byte(v) {
				return strs[j][:i]
			}
		}

	}
	return short
}

func short(strs []string)string{
	 min:=len(strs[0])
	shortt:=strs[0]
	for k,v:=range strs{
		if len(strs[k])<min{
			min=len(strs[k])
			shortt=v
			}
		
	}
	return shortt
}


func longestCommonPrefix(strs []string) string {
	short:=short(strs)
	for k,v:=range short{
		for i:=;i<len(strs);i++{
			if strs[i][k]!=byte(v){
				return strs[i][:k]
			}
		}
	}
	return short

}


