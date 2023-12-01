package main

import (
	"fmt"
	"strings"
)

// 加密函数
func Encrypt(plaintext string, key [][]int) string {
	// 将明文转换为大写字母，并删除空格
	plaintext = sanitizeText(plaintext)

	// 对明文进行分组
	blocks := makeBlocks(plaintext, len(key))

	// 进行加密计算
	ciphertext := ""
	for _, block := range blocks {
		encryptedBlock := encryptBlock(block, key)
		ciphertext += encryptedBlock
	}

	return ciphertext
}

// 解密函数
func Decrypt(ciphertext string, key [][]int) string {
	// 将密文转换为大写字母，并删除空格
	ciphertext = sanitizeText(ciphertext)

	// 对密文进行分组
	blocks := makeBlocks(ciphertext, len(key))

	// 进行解密计算
	plaintext := ""
	for _, block := range blocks {
		decryptedBlock := decryptBlock(block, key)
		plaintext += decryptedBlock
	}

	return plaintext
}

// 对文本进行预处理，将其转换为大写字母并删除空格
func sanitizeText(text string) string {
	text = strings.ToUpper(text)
	text = strings.ReplaceAll(text, " ", "")
	return text
}

// 将文本分成块
func makeBlocks(text string, blockSize int) []string {
	blocks := []string{}
	for i := 0; i < len(text); i += blockSize {
		if i+blockSize <= len(text) {
			blocks = append(blocks, text[i:i+blockSize])
		} else {
			// 如果最后一个分组不够 blockSize 长度，则使用填充字符填充
			block := text[i:]
			padding := blockSize - len(block)
			for j := 0; j < padding; j++ {
				block += "X"
			}
			blocks = append(blocks, block)
		}
	}
	return blocks
}

// 加密单个分组
func encryptBlock(block string, key [][]int) string {
	encryptedBlock := ""
	for _, row := range key {
		sum := 0
		for i, ch := range block {
			num := int(ch - 'A')
			sum += row[i] * num
		}
		encryptedChar := byte((sum % 26) + 'A')
		encryptedBlock += string(encryptedChar)
	}
	return encryptedBlock
}

// 解密单个分组
func decryptBlock(block string, key [][]int) string {
	det := determinant(key)
	invKey := inverseKey(key, det)

	decryptedBlock := ""
	for _, row := range invKey {
		sum := 0
		for i, ch := range block {
			num := int(ch - 'A')
			sum += row[i] * num
		}
		decryptedChar := byte((sum % 26) + 'A')
		decryptedBlock += string(decryptedChar)
	}
	return decryptedBlock
}

// 计算矩阵的行列式
func determinant(matrix [][]int) int {
	det := matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	return det
}

// 求矩阵的逆矩阵
func inverseKey(key [][]int, det int) [][]int {
	invDet := 0
	for i := 0; i < 26; i++ {
		if (det*i)%26 == 1 {
			invDet = i
			break
		}
	}

	// 计算伴随矩阵
	adjMatrix := [][]int{
		{key[1][1], -key[0][1]},
		{-key[1][0], key[0][0]},
	}

	// 计算逆矩阵
	invKey := [][]int{}
	for _, row := range adjMatrix {
		invRow := []int{}
		for _, num := range row {
			invNum := (num * invDet) % 26
			if invNum < 0 {
				invNum += 26
			}
			invRow = append(invRow, invNum)
		}
		invKey = append(invKey, invRow)
	}

	return invKey
}

func main() {
	key := [][]int{{6, 24, 1}, {13, 16, 10}, {20, 17, 15}}
	plaintext := "HELLO"

	ciphertext := Encrypt(plaintext, key)
	fmt.Println("Ciphertext:", ciphertext)

	decrypted := Decrypt(ciphertext, key)
	fmt.Println("Decrypted:", decrypted)
}
