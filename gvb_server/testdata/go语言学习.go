// package main
//
// import "fmt"
//
//	type a struct {
//		i, j int
//	}
//
//	func functin(A *a) *a {
//		fmt.Printf("%p", &A)
//
//		fmt.Println()
//		return A
//	}
//
//	func main() {
//		aa := &a{1, 2}
//		bb := functin(aa)
//
//		fmt.Printf("%p,%p", &aa, &bb)
//	}

//package main
//
//import "fmt"
//
//func main() {
//	s := make([]int, 0, 10)
//	appendData(s, 1, 2, 3, 4, 5, 6)
//	fmt.Println(s)
//}
//
//func appendData[T comparable](s []T, data ...T) {
//	s = append(s, data...)
//}

package main

import "fmt"

func main() {
	s := make([]int16, 0, 500)
	fmt.Println(cap(s))
	for i := 0; i < 513; i++ {
		s = append(s, int16(i))
	}
	fmt.Println(cap(s))
}
