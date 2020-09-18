package main

func firstUniqChar(s string) byte {
	var res [26]int
	for i:=0;i<len(s);i++{
		res[s[i]-'a']++
	}
	for i:=0;i<len(s);i++{
		if res[s[i]-'a']==1{
			return s[i]
		}
	}
return ' '
}
