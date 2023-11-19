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
		// 应用加法密码加密公式
		var encryptedChar int
		if unicode.IsUpper(ch) {
			encryptedChar = (chASCII+k)%26 + 65
		} else if unicode.IsLower(ch) {
			encryptedChar = (chASCII+k)%26 + 97
		}
		// 将加密后的ASCII值转换回字符
		encryptedCh := rune(encryptedChar)
		// 将加密后的字符追加到密文中
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
			decryptedChar = (chASCII-k+26)%26 + 65
		} else if unicode.IsLower(ch) {
			decryptedChar = (chASCII-k+26)%26 + 97
		}
		// 将解密后的ASCII值转换回字符
		decryptedCh := rune(decryptedChar)
		// 将解密后的字符追加到明文中
		plaintext += string(decryptedCh)
	}
	return plaintext
}

func main() {
	plaintext := "HellO"
	key := 3
	ciphertext := encrypt(plaintext, key)
	fmt.Println("密文:", ciphertext)
	decryptedText := decrypt(ciphertext, key)
	fmt.Println("解密后的文本:", decryptedText)
}

/*
这段代码定义了两个函数：Encrypt和Decrypt，分别用于加密和解密加法密码。
Encrypt函数接收明文消息和密钥k，返回密文。
Decrypt函数接收密文和相同的密钥k，返回解密后的明文。
主函数通过使用密钥为3将明文"HELLO"加密，并将密文解密回原始明文进行演示。
*/
