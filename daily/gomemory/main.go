package main

import "sync"

func main(){
	lock1:=sync.Mutex{}
	lock2:=sync.Mutex{}
	go a()
	go b()
}
func a(){
	lock1.lock
	defer lock1.unlock
	i++
	lock2.lock
}
func b(){
	lock2.lock
	j++
	lock1.lock

}