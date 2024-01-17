package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("加载中...")
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r加载中%c", r)
			time.Sleep(delay)
		}

	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func A() {
//	// A函数的执行逻辑
//	// 创建一个通道用于接收信号
//	done := make(chan bool)
//
//	// 启动B函数，并传递通道作为参数
//	go func() {
//		B(done)
//	}()
//
//	time.Sleep(2 * time.Second)
//	// A函数结束时发送信号给通道
//	done <- true
//	fmt.Println("A函数结束")
//}
//
//func B(done chan bool) {
//	// B函数的死循环逻辑
//	for {
//		select {
//		case <-done: // 接收到信号，退出循环
//			return
//		default:
//			// 执行B函数的逻辑
//			for _, r := range `-\|/` {
//				fmt.Printf("\r回答中%c", r)
//				time.Sleep(100 * time.Millisecond)
//			}
//		}
//	}
//}
//
//func main() {
//	A() // 调用A函数
//}
