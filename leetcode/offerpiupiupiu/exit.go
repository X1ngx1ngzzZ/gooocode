package main

import (
	"fmt"
)

func main() {
	a := [][]int{{1, 1, 1}, {2, 2, 2}}
	fmt.Println(len(a))
	fmt.Println(len(a[0]))
}

func exist(board [][]byte, word string) bool {
	//转成字符
	words := []byte(word)
	//可以提前存一下长度，但是要传进去就比较麻烦所以不存了
	//.	m:=len(board)
	//	n:=len(board[0])
	for i, value1 := range board {
		//用board[i][j]就可以表示每一个字符了，不用这个内循环的值
		for j, _ := range value1 {
			//这个0代表字符中的第一个元素，从第一个开始找所以初始值设为0
			if deal(board, words, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func deal(board [][]byte, words []byte, i, j, k int) bool {
	//把判断false放在前面，减少浪费
	if i > len(board)-1 || i < 0 || j > len(board[0])-1 || j < 0 || board[i][j] != words[k] {
		return false
	}
	if k == len(words)-1 {
		return true
	}
	//值相等的时候再开始比较下一个
	//拿个值暂存一下，然后设置为/，避免找到之前的
	tmp := board[i][j]
	board[i][j] = '/'
	//同时看上下左右是否符合
	res := deal(board, words, i+1, j, k+1) || deal(board, words, i, j+1, k+1) || deal(board, words, i-1, j, k+1) || deal(board, words, i, j-1, k+1)
	board[i][j] = tmp
	return res
}
