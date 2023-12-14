package main

import "fmt"

// 型
// byte(unit8)型

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA) // [72 73]

	// 文字列に変換
	fmt.Println(string(byteA)) // HI

	// 文字列をbyte型に変換
	c := []byte("HI")
	fmt.Println(c) // [72 73]
}
