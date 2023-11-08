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

// 2. 编程实现利用扩展欧几里德算法求解最大公因子和乘法逆元
func extendedGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

func modInverse(a, m int) int {
	gcd, x, _ := extendedGCD(a, m)
	if gcd != 1 {
		// 逆元不存在
		return -1
	}
	// 确保逆元是正数
	result := (x%m + m) % m
	return result
}

//3. 实现单表代换密码中的加法密码

func main() {
	//1. 编程实现模n的快速指数运算
	//30的37次方模77，答案为2
	fmt.Println(qmi(30, 37, 77))

	//2. 编程实现利用扩展欧几里德算法求解最大公因子和乘法逆元
	//求最大公因子
	x := 10
	y := 11
	gcd, _, _ := extendedGCD(x, y)
	fmt.Printf("%d 和 %d 的最大公因子为:%d\n", x, y, gcd)

	//求乘法逆元
	a := 7  // 要找逆元的数
	m := 11 // 模数
	inverse := modInverse(a, m)
	if inverse == -1 {
		fmt.Printf("%d 没有乘法逆元模 %d\n", a, m)
	} else {
		fmt.Printf("乘法逆元: %d\n", inverse)
	}
}
