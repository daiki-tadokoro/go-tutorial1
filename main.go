package main

import "fmt"

// 型
// 整数型

func main() {
	var i int = 100 // 環境依存

	// var i2 int8 = 100 // 環境依存せずに8bitの整数型を指定できる

	fmt.Println(i + 50) // 100 + 50 = 150

	var i3 int32 = 200
	// 型の調べ方
	fmt.Printf("%T\n", i3) // int32

	fmt.Printf("%T\n", int(i3)) // int32からintに変換
	fmt.Println(i + int(i3))    // 300
}
