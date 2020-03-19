func myAtoi(str string) int {
    str = strings.TrimSpace(str)
    if len(str) == 0 {
        return 0
    }
    
    num := 0
    flag := false
    if str[0] =='-' || str[0] == '+'{ 
        if str[0] == '-' {
            flag = true
        }
        str = str[1:]
    }

    for k,v := range str {
        if str[k] < 48 || str[k] > 57 {
            break
        }
        c := int(v-48)
        num = num*10 + c
        if num > (1<<31-1) || num < -1<<31 {
            if flag == true {
                return -1<<31
            } else {
                return (1<<31-1)
            } 
        }
    }
    fmt.Println(num)
    if flag == true {
        num = -num
    }
    return num
}