package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 生成密钥
func generateKey(message, key string) string {
	x := len(message) / len(key)
	y := len(message) % len(key)
	key = strings.Repeat(key, x) + key[:y] //重复key x次，然后加上key从0到y的值，正好和message长度一致
	return key
}

func encrypt_vg(message, key string) string {
	key = generateKey(message, key)
	var cipherText string
	for i := range message {
		if message[i] != ' ' {
			offset := 'A'
			if unicode.IsLower(rune(message[i])) {
				offset = 'a'
			}

			m := int(message[i]) - int(offset)
			k := int(unicode.ToUpper(rune(key[i])) - 'A')
			x := (m + k) % 26
			x += int(offset)
			cipherText += string(x)

		} else {
			cipherText += " "
		}
	}
	return cipherText
}

// decrypt 函数用于进行维吉尼亚密码的解密
func decrypt_vg(cipherText, key string) string {
	key = generateKey(cipherText, key)
	var originalText string
	for i := range cipherText {
		if cipherText[i] != ' ' {
			offset := 'A'
			if unicode.IsLower(rune(cipherText[i])) {
				offset = 'a'
			}

			m := int(cipherText[i]) - int(offset)
			k := int(unicode.ToUpper(rune(key[i])) - 'A')
			x := (m - k + 26) % 26
			x += int(offset)
			originalText += string(x)

		} else {
			originalText += " "
		}
	}
	return originalText
}

func main() {
	message := "Hello World"
	key := "KEY"
	fmt.Println("明文为：", message)
	cipherText := encrypt_vg(message, key)
	fmt.Println("密文为：", cipherText)
	originalText := decrypt_vg(cipherText, key)
	fmt.Println("解密后为：", originalText)
}
