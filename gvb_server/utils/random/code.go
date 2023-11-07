package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//func Code(len int) string {
//	rand.Seed(time.Now().UnixNano()) //指定种子
//
//	return fmt.Sprintf("%4v", rand.Intn(10000))
//}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var code strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&code, "%d", numeric[rand.Intn(r)])
	}
	return code.String()
}
