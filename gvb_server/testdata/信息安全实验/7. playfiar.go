package main

import (
	"fmt"
	"strings"
)

type Martix struct {
	ma        [5][5]rune    //矩阵
	uniqueMap map[rune]bool //去重哈希
	key       string        //密钥
	posJ      []int         //存入j的坐标
	posX      []int         //存入插入x的坐标
}

func (m *Martix) New() *Martix {
	m.uniqueMap = make(map[rune]bool) //初始化
	m.ma = [5][5]rune{}               //初始化
	m.key = "HELLO"                   //暂时先赋值
	//scanner := bufio.NewScanner(os.Stdin)
	//fmt.Printf("Please enter the key: ")
	//scanner.Scan()
	//input := scanner.Text()
	//m.key = strings.ToUpper(input) //转换为大写

	return m
}

func (m *Martix) NewMatrix() error {
	var pos int
	alphabet := "ABCDEFGHIKLMNOPQRSTUVWXYZ"
	fill := m.key + alphabet
	fill = strings.ReplaceAll(fill, "J", "I")

	//填写矩阵
	for _, str := range fill {
		if m.uniqueMap[str] {
			continue
		}

		m.ma[pos/5][pos%5] = str
		pos++
		m.uniqueMap[str] = true
	}

	//将矩阵输出
	fmt.Println("The matrix is: ")
	for _, runes := range m.ma {
		for _, ch := range runes {
			fmt.Printf("%c ", ch)
		}
		fmt.Println()
	}

	return nil
}

func (m *Martix) encryption(plaintext string) (ciphertext string) {
	//将明文进行处理
	plaintext = strings.ToUpper(plaintext)
	for i, ch := range plaintext { //记录J的下标
		if ch == 'J' {
			m.posJ = append(m.posJ, i)
		}
	}
	plaintext = strings.ReplaceAll(plaintext, "J", "I") //将J全部置换为I
	for i := 0; i < len(plaintext)-1; i += 2 {
		if plaintext[i] == plaintext[i+1] {
			plaintext = insertChar(i, plaintext, "X") //在两者中间插入一个x
			m.posX = append(m.posX, i+1)
		}
	}
	if len(plaintext)%2 == 1 { //约定好的字母是x
		plaintext += "X"
		m.posX = append(m.posX, len(plaintext)-1)
	}

	//加密
	for i := 0; i < len(plaintext)-1; i += 2 {
		row1, col1 := findPosition(m.ma, rune(plaintext[i]))
		row2, col2 := findPosition(m.ma, rune(plaintext[i+1]))

		var encryptedPair string
		if row1 == row2 { // 同一行
			encryptedPair = string(m.ma[row1][(col1+1)%5]) + string(m.ma[row2][(col2+1)%5])
		} else if col1 == col2 { // 同一列
			encryptedPair = string(m.ma[(row1+1)%5][col1]) + string(m.ma[(row2+1)%5][col2])
		} else { // 不同行不同列
			encryptedPair = string(m.ma[row1][col2]) + string(m.ma[row2][col1])
		}

		ciphertext += encryptedPair
	}

	return
}

func (m *Martix) decryption(ciphertext string) (plaintext string) {
	//将明文进行处理
	ciphertext = strings.ToUpper(ciphertext)

	for i := 0; i < len(ciphertext)-1; i += 2 {
		row1, col1 := findPosition(m.ma, rune(ciphertext[i]))
		row2, col2 := findPosition(m.ma, rune(ciphertext[i+1]))

		//解密操作
		var decryptedPair string
		if row1 == row2 { // 同一行
			decryptedPair = string(m.ma[row1][(col1+4)%5]) + string(m.ma[row2][(col2+4)%5])
		} else if col1 == col2 { // 同一列
			decryptedPair = string(m.ma[(row1+4)%5][col1]) + string(m.ma[(row2+4)%5][col2])
		} else { // 不同行不同列
			decryptedPair = string(m.ma[row1][col2]) + string(m.ma[row2][col1])
		}

		plaintext += decryptedPair
	}

	//整理解密后的明文
	plaintextByte := []byte(plaintext)
	for _, cnt := range m.posJ {
		plaintextByte[cnt] = 'J'
	}
	// 去除插入的'X'
	for _, cnt := range m.posX {
		if cnt == 0 {
			plaintextByte = plaintextByte[1:]
		} else {
			plaintextByte = append(plaintextByte[:cnt-1], plaintextByte[cnt:]...)
		}
	}

	plaintext = string(plaintextByte)

	//如果是循环输入，那么需要清空posX和posY
	m.posJ = make([]int, 0)
	m.posX = make([]int, 0)

	return
}

func main() {
	var mar Martix
	mar.New().NewMatrix() //初始化矩阵

	miwen := mar.encryption("HELLOWORLD")
	fmt.Println("密文：", miwen)

	minwen := mar.decryption(miwen)
	fmt.Println("明文：", minwen)
}

func insertChar(idx int, str, ch string) string {
	return str[:idx] + ch + str[idx:]
}

func findPosition(matrix [5][5]rune, ch rune) (int, int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrix[i][j] == ch {
				return i, j
			}
		}
	}
	return -1, -1
}
