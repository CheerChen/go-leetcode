package gofeature

import (
	"fmt"
	"time"
)

func testPanic() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		for {
			time.Sleep(time.Second)
			// panic会让当前goroutine退出，所以每一秒创建一个新的gorutine
			go func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println("panic", err)
					}
				}()
				proc()
			}()
		}

	}()

	select {}
}

func proc() {
	panic("ok")
}
