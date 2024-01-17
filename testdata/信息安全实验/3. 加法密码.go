package main

import (
	"fmt"
	"unicode"
)

// Encrypt 使用加法密码将明文加密，使用密钥k
func encrypt(plaintext string, k int) string {
	ciphertext := ""
	for _, ch := range plaintext {
		// 将字符转换为ASCII值
		chASCII := int(ch)
		// 加法密码加密
		var encryptedChar int
		if unicode.IsUpper(ch) {
			encryptedChar = ((chASCII-65)+k)%26 + 65
		} else if unicode.IsLower(ch) {
			encryptedChar = ((chASCII-97)+k)%26 + 97
		}
		encryptedCh := rune(encryptedChar)
		ciphertext += string(encryptedCh)
	}
	return ciphertext
}

// Decrypt 使用加法密码将密文解密，使用密钥k
func decrypt(ciphertext string, k int) string {
	plaintext := ""
	for _, ch := range ciphertext {
		// 将字符转换为ASCII值
		chASCII := int(ch)
		// 应用加法密码解密公式
		var decryptedChar int
		if unicode.IsUpper(ch) {
			decryptedChar = ((chASCII-65)-k+26)%26 + 65
		} else if unicode.IsLower(ch) {
			decryptedChar = ((chASCII-97)-k+26)%26 + 97
		}
		decryptedCh := rune(decryptedChar)
		plaintext += string(decryptedCh)
	}
	return plaintext
}

func main() {
	plaintext := "Hello"
	key := 3
	fmt.Println("明文为：", plaintext)
	ciphertext := encrypt(plaintext, key)
	fmt.Println("密文为：", ciphertext)
	decryptedText := decrypt(ciphertext, key)
	fmt.Println("解密后为：", decryptedText)
}
