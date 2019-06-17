package main

import (
	"fmt"
	"sync"
	"time"
)

func write(ch chan int, wait *sync.WaitGroup) {
	defer wait.Done()
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("put data:", i)
	}
	close(ch)
}

func read(ch chan int, wait *sync.WaitGroup) {
	defer wait.Done()
	for {
		// var b int
		b, ok := <-ch
		fmt.Println(b)
		if !ok {
			// os.Exit(1)
			break
		}
		time.Sleep(time.Second)
	}
}

func main() {
	//WaitGroup的用途是使得主线程一直阻塞等待直到所有相关的子goroutine都已经完成了任务。
	//sync.WaitGroup只有3个API
	//
	//Add()      # 添加计数
	//Done()    # 减掉计数，等价于Add(-1)，这样sync.WaitGroup只有两个API了
	//Wait()      # 阻塞直到计数为零

	wg := sync.WaitGroup{}
	wg.Add(2)
	intChan := make(chan int, 5)
	go write(intChan, &wg)
	go read(intChan, &wg)

	wg.Wait()
	// time.Sleep(20 * time.Second)
}

