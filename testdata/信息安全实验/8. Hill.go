package main

import (
	"fmt"
	"strings"
)

type Hill struct {
}

// 加密函数
func (Hill) encrypt(plaintext string, key [][]int) string {
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
func (Hill) decrypt(ciphertext string, invKey [][]int) string {
	// 将密文转换为大写字母，并删除空格
	ciphertext = sanitizeText(ciphertext)

	// 对密文进行分组
	blocks := makeBlocks(ciphertext, len(invKey))

	// 进行解密计算
	plaintext := ""
	for _, block := range blocks {
		decryptedBlock := decryptBlock(block, invKey)
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
	blocks := make([]string, 0)
	for i := 0; i < len(text); i += blockSize {
		if blockSize+i <= len(text) {
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
func decryptBlock(block string, invkey [][]int) string {
	decryptedBlock := ""
	for _, row := range invkey {
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

/*
是的，Hill密码算法的密钥矩阵必须是方阵。这是因为Hill密码算法的加密和解密过程都涉及到矩阵乘法和求逆。
在这个过程中，只有方阵才能保证运算的正确性。此外，这个密钥矩阵还必须是可逆的，也就是说它的行列式必须和26互质。
如果密钥矩阵不是可逆的，那么就无法进行解密。
*/

// 计算矩阵的行列式
//func determinant(matrix [][]int) (int, error) {
//	dim := len(matrix)
//	if dim == 0 {
//		return 0, errors.New("matrix is empty")
//	}
//	for i := range matrix {
//		if len(matrix[i]) != dim {
//			return 0, errors.New("matrix is not square")
//		}
//	}
//	if dim == 1 {
//		return matrix[0][0], nil
//	}
//	if dim == 2 {
//		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0], nil
//	}
//	var det int
//	for i := range matrix {
//		submatrix := make([][]int, dim-1)
//		for j := range submatrix {
//			submatrix[j] = make([]int, dim-1)
//		}
//		for j := 1; j < dim; j++ {
//			for k := range matrix[j] {
//				if k < i {
//					submatrix[j-1][k] = matrix[j][k]
//				} else if k > i {
//					submatrix[j-1][k-1] = matrix[j][k]
//				}
//			}
//		}
//		subDet, err := determinant(submatrix)
//		if err != nil {
//			return 0, err
//		}
//		if i%2 == 0 {
//			det += matrix[0][i] * subDet
//		} else {
//			det -= matrix[0][i] * subDet
//		}
//	}
//	return det, nil
//}

// 求矩阵的逆矩阵
//func inverseKey(matrix [][]float64) [][]float64 {
//	n := len(matrix)
//	inverse := make([][]float64, n)
//	for i := range inverse {
//		inverse[i] = make([]float64, n)
//		inverse[i][i] = 1
//	}
//
//	for i := 0; i < n; i++ {
//		for j := i; j < n; j++ {
//			if matrix[j][i] != 0 {
//				matrix[i], matrix[j] = matrix[j], matrix[i]
//				inverse[i], inverse[j] = inverse[j], inverse[i]
//				break
//			}
//		}
//
//		for j := 0; j < n; j++ {
//			if i != j {
//				ratio := matrix[j][i] / matrix[i][i]
//				for k := 0; k < n; k++ {
//					matrix[j][k] -= ratio * matrix[i][k]
//					inverse[j][k] -= ratio * inverse[i][k]
//				}
//			}
//		}
//	}
//
//	for i := 0; i < n; i++ {
//		a := matrix[i][i]
//		for j := 0; j < n; j++ {
//			matrix[i][j] /= a
//			inverse[i][j] /= a
//		}
//	}
//
//	return inverse
//}

func main() {
	var h Hill
	key := [][]int{
		{17, 17, 5},
		{21, 18, 21},
		{2, 2, 19}}

	invKey := [][]int{
		{4, 9, 15},
		{15, 17, 6},
		{24, 0, 17}}

	plaintext := "HELLOX"

	ciphertext := h.encrypt(plaintext, key)
	fmt.Println("Ciphertext:", ciphertext)

	decrypted := h.decrypt(ciphertext, invKey)
	fmt.Println("Decrypted:", decrypted)
}
