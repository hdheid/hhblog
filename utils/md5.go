package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 加密
func Md5(src []byte) string {
	//在给定的代码中，m := md5.New() 和 m.Write(src) 用于计算给定数据 src 的 MD5 哈希值。
	//
	//md5.New() 创建一个新的 MD5 哈希对象 m，用于计算 MD5 哈希值。
	//m.Write(src) 将数据 src 写入到哈希对象 m 中，以便计算哈希值。src 可以是任何字节切片或字符串。
	//通过上述代码的执行，可以得到 src 数据的 MD5 哈希值。要获取最终的哈希值，可以调用 m.Sum(nil) 方法，它会返回一个字节切片表示的哈希值。如果需要将哈希值表示为十六进制字符串，可以使用 hex.EncodeToString(m.Sum(nil))。
	m := md5.New()
	m.Write(src)
	//m.Sum(nil) 调用哈希对象的 Sum() 方法，返回哈希值的字节切片。
	//hex.EncodeToString() 函数将字节切片转换为十六进制格式的字符串表示。
	//hash 是一个字符串变量，用于存储转换后的十六进制字符串值。
	//综上所述，hash := hex.EncodeToString(m.Sum(nil)) 的目的是计算哈希值并将其转换为十六进制字符串，然后将结果存储在变量 hash 中。
	hash := hex.EncodeToString(m.Sum(nil))
	return hash
}
