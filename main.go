package main

import "fmt"

// 型
// インターフェース型

func main() {
	var x interface{} // 初期値はnil
	fmt.Println(x)    // <nil> あらゆる型を入れることができる

	x = 1
	fmt.Println(x) // 1

	x = 3.14
	fmt.Println(x) // 3.14

	x = [3]int{1, 2, 3}
	fmt.Println(x) // [1 2 3]

	x = 2
	// fmt.Println(x + 3)
	// 参照はできるがデータ特有の演算はできない
}
