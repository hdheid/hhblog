package main

import "fmt"

// 1. 编程实现模n的快速指数运算
func qmi(m, e, n int64) (ans int64) {
	ans = 1
	for e != 0 {
		if e&1 == 1 {
			ans = ans * m % n
		}
		e >>= 1
		m = m * m % n
	}
	return
}

func main() {
	//1. 编程实现模n的快速指数运算
	//30的37次方模77，答案为2
	fmt.Println("30的37次方模77，答案为：", qmi(30, 37, 77))
}
