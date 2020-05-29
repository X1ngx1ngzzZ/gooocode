package main

import "fmt"

func main() {
	m1 := []int{1, 2, 3, 4, 5}
	m2 := []int{2, 3}
	fmt.Println(Complement(m1, m2))
}

func Complement(A []int, B []int) []int {
	// wriifte code here
	if B == nil {
		return A
	} else if A == nil && B == nil {
		return A
	} else if A == nil && B != nil {
		return nil
	}
	m := make(map[int]bool)
	//	fmt.Println(A, B, m)
	var res []int
	for _, v := range B {
		m[v] = true
	}
	//	fmt.Println(m)
	for _, v := range A {
		if m[v] {

		} else {
			res = append(res, v)
		}
	}
	return res
}
