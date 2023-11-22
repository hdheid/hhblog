package main

import (
	"fmt"
	"unicode"
)

func extendedGCD_aff(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD_aff(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

func modInverse_aff(a, m int) int {
	g, x, _ := extendedGCD_aff(a, m)
	if g != 1 {
		// 逆元不存在
		return -1
	}
	// 确保逆元是正数
	result := (x%m + m) % m
	return result
}

func encrypt_aff(msg string, a, b int) string {
	var res string
	for _, r := range msg {
		if r != ' ' {
			if unicode.IsUpper(r) {
				res += string((a*int(r-'A')+b)%26 + 'A')
			} else if unicode.IsLower(r) {
				res += string((a*int(r-'a')+b)%26 + 'a')
			}
		} else {
			res += " "
		}
	}
	return res
}

func decrypt_aff(cipher string, a, b int) string {
	var res string
	aInv := modInverse_aff(a, 26)
	for _, r := range cipher {
		if r != ' ' {
			if unicode.IsUpper(r) {
				res += string(aInv*(int(r-'A')-b+26)%26 + 'A')
			} else if unicode.IsLower(r) {
				res += string(aInv*(int(r-'a')-b+26)%26 + 'a')
			}
		} else {
			res += " "
		}
	}
	return res
}

func main() {
	message := "HELLO world"
	a, b := 6, 8
	gcd, _, _ := extendedGCD_aff(a, 26)
	if gcd != 1 {
		fmt.Println("Error: a and 26 are not coprime")
		return
	}
	fmt.Println("明文为：", message)
	cipher := encrypt_aff(message, a, b)
	fmt.Println("密文为：", cipher)
	decipher := decrypt_aff(cipher, a, b)
	fmt.Println("解密后为：", decipher)
}