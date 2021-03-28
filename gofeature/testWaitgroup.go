package gofeature

import (
	"fmt"
	"sync"
	"time"
)

func testWaitgroup() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	waitCh:=make(chan struct{})
	go func() {
		defer close(waitCh)
		wg.Wait()
	}()
	select {
	case <-waitCh:
		break
	case <-time.After(timeout):
		return true
	}
	return false
}