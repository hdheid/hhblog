package main

import (
	"fmt"
	"sort"
	"strings"
)

// 列置换密码加密
func columnarTranspositionEncrypt(input, key string) string {
	input = strings.ReplaceAll(input, " ", "")
	columns := [][]rune{}
	mp := map[rune]int{}
	for i, k := range key {
		mp[k] = i
	}

	for i := 0; i < len(input); i += len(key) {
		var col []rune
		for j := i; j < len(input) && j < i+len(key); j++ {
			col = append(col, rune(input[j]))
		}
		columns = append(columns, col)
	}

	sortKey := []rune(key)
	sort.Slice(sortKey, func(i, j int) bool {
		return sortKey[i] < sortKey[j]
	})

	var result strings.Builder
	for _, r := range sortKey {
		idx := mp[r]
		var res string
		for _, column := range columns {
			if idx < len(column) {
				res += string(column[idx])
			}
		}
		result.WriteString(res)
	}

	return result.String()
}

// 列置换密码解密
func columnarTranspositionDecrypt(input, key string) string {
	columns := make([][]rune, len(key))
	mp := map[rune]int{}
	for i, k := range key {
		mp[k] = i
	}

	sortKey := []rune(key)
	sort.Slice(sortKey, func(i, j int) bool {
		return sortKey[i] < sortKey[j]
	})

	colLength := len(input) / len(key)
	extraChars := len(input) % len(key)

	startIdx := 0
	for _, r := range sortKey {
		idx := mp[r]
		length := colLength
		if idx < extraChars {
			length++
		}
		columns[idx] = []rune(input[startIdx : startIdx+length])
		startIdx += length
	}

	var result strings.Builder
	for i := 0; i < colLength; i++ {
		for _, col := range columns {
			if i < len(col) {
				result.WriteRune(col[i])
			}
		}
	}

	for _, col := range columns {
		if colLength < len(col) {
			result.WriteRune(col[colLength])
		}
	}

	return result.String()
}

func main() {
	// 测试列置换密码
	key := "CBA"
	fmt.Println("\n列置换密码:")
	plainText := "hello world"
	cipherText := columnarTranspositionEncrypt(plainText, key)
	fmt.Println("明文:", plainText)
	fmt.Println("密文:", cipherText)
	fmt.Println("解密:", columnarTranspositionDecrypt(cipherText, key))
}
