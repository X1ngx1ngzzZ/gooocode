package main

import (
	"fmt"
	"sync"
)

var x int

//建这个等待队列并不能改变临界区的问题，只是为了控制并发数量
//使用这个来实现并发任务的同步，内部维护着一个计数器，计数器
//的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计
//数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。
//通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。
var wg sync.WaitGroup

//建个互斥锁
//互斥锁完全互斥，但有时候获取资源的时候是没必要完全加锁的
var lock sync.Mutex



func main() {
	fmt.Println(x)
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)

}

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
