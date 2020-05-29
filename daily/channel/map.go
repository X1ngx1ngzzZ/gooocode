package main

import (
	"fmt"
	"sync"
)

var (
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)
var m = make(map[string]int)
var sm = sync.Map{}

//map不是并发安全的，对传统map并发读写一多就会panic
/*
func main() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			//转成字符串
			key := strconv.Itoa(n)
			//设置k,v
			rwlock.Lock()
			set(key, n)
			fmt.Printf("k=%v,v=%v\n", key, getkey(key))
			rwlock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
*/
func set(k string, v int) {
	m[k] = v
}

func getkey(k string) int {
	return m[k]
}

//GO提供了可并发的map

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			sm.Store(n, n+100)
			v, _ := sm.Load(n)
			fmt.Printf("k=%v,v=%v\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
