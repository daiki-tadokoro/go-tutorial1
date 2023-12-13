package main

import "fmt"

// 型
// 文字列型

func main() {
	var s string = "Hello Golang"
	fmt.Printf("%T\n", s) // string

	var s1 string = "300"
	fmt.Printf("%T\n", s1) // string

	// 複数行の文字列
	var s2 string = `Hello
	Golang
	!`
	fmt.Println(s2)
	fmt.Printf("%T\n", s2) // string

	// ダブルクォーテーションを文字列として扱う
	var s3 string = "\""
	fmt.Println(s3) // "

	var s4 string = `"`
	fmt.Println(s4) // "

	// 文字列の長さ
	fmt.Println(s[0])         // 72
	fmt.Println(string(s[0])) // H

}
