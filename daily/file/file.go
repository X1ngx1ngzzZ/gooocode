package main

import (
	"io"
	"os"
)

func main() {
	path1 := "./abc.txt"
	path2 := "./cde.txt"
	f1, _ := os.Open(path1)
	f2, _ := os.OpenFile(path2, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer f1.Close()
	defer f2.Close()
	io.Copy(f2, f1)
}
