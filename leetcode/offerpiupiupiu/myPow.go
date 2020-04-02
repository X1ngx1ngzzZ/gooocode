func myPow(x float64, n int) float64 {
	switch n {
	case 0: return 1
	case 1: return x
	case -1: return 1/x
	}
	a:=myPow(x,n/2)
	b:=myPow(x,n%2)
	return a*a*b
}