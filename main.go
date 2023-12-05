package main

import "fmt"

func main() {
	// 明示的な定義
	// var 変数名 型 = 値
	var i int = 100

	fmt.Println(i) // 100

	var s string = "Hello World"
	fmt.Println(s) // Hello World

	var t, f bool = true, false
	fmt.Println(t, f) // true false

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2) // 200 Golang

	var i3 int
	var s3 string

	fmt.Println(i3, s3) // 0

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3) // 300 Go
}
