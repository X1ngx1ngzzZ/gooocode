package main

import (
	"fmt"
	"os"
)

func main() {
	pid := os.Getpid()
	ppid := os.Getppid()
	fmt.Printf("%d\n", pid)
	fmt.Printf("%d\n", ppid)
}
