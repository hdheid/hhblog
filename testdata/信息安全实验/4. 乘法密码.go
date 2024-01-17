package main

import (
	"errors"
	"fmt"
	"unicode"
)

//a和A暂时没办法加密

// Encrypt 使用乘法密码算法对明文进行加密，使用密钥 k 和模数 n
func encrypt_mul(plaintext string, k int) string {
	ciphertext := ""
	for _, ch := range plaintext {
		// 将字符转换为ASCII值
		chASCII := int(ch)
		// 应用乘法密码加密公式
		var encryptedChar int
		if unicode.IsUpper(ch) {
			chASCII -= 65
			encryptedChar = (chASCII*k)%26 + 65
		} else if unicode.IsLower(ch) {
			chASCII -= 97
			encryptedChar = (chASCII*k)%26 + 97
		}
		encryptedCh := rune(encryptedChar)
		ciphertext += string(encryptedCh)
	}
	return ciphertext
}

// Decrypt 使用乘法密码算法对密文进行解密，使用密钥 k 和模数 n
func decrypt_mul(ciphertext string, k int) string {
	plaintext := ""
	// 计算 k 的模逆元
	kInverse, err := modInverse_mul(k, 26)
	if err != nil {
		panic(err)
	}

	for _, ch := range ciphertext {
		// 将字符转换为ASCII值
		chASCII := int(ch)
		// 应用乘法密码解密公式
		var encryptedChar int
		if unicode.IsUpper(ch) {
			chASCII -= 65
			encryptedChar = (chASCII*kInverse)%26 + 65
		} else if unicode.IsLower(ch) {
			chASCII -= 97
			encryptedChar = (chASCII*kInverse)%26 + 97
		}
		encryptedCh := rune(encryptedChar)
		plaintext += string(encryptedCh)
	}
	return plaintext
}

// 扩展欧几里得算法
func extendedGCD_mul(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD_mul(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

func modInverse_mul(a, m int) (int, error) {
	gcd, x, _ := extendedGCD_mul(a, m)
	if gcd != 1 {
		// 逆元不存在
		return -1, errors.New("错误：给定的密钥和模数不互质。")
	}
	// 确保逆元是正数
	result := (x%m + m) % m
	return result, nil
}

func main() {
	plaintext := "Hello"
	key := 5
	fmt.Println("明文为：", plaintext)
	ciphertext := encrypt_mul(plaintext, key)
	fmt.Println("密文为：", ciphertext)
	decryptedText := decrypt_mul(ciphertext, key)
	fmt.Println("解密后为：", decryptedText)
}
