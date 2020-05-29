func ReversePolishNotation1(input []string) int {
    number := []int{}
	for _, val := range input{
		l := len(number)
		switch val {
		case "+":
			number  = append(number[:l -2], number[l-2] + number[l-1])
		case "-":
			number  = append(number[:l -2], number[l-2] - number[l-1])
		case "*":
			number  = append(number[:l -2], number[l-2] * number[l-1])
		case "/":
			number  = append(number[:l -2], number[l-2] / number[l-1])
		default:
			num, _ := strconv.Atoi(val)
			number  = append(number, num)
		}
	}
	return number[0]
}
