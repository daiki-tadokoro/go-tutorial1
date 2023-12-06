package main

import "fmt"

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	// 暗黙的な定義(型推論)
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4) // 400

	// 値の更新
	i4 = 450
	fmt.Println(i4) // 450

	// 再定義はできない
	// i4 := 500
	// fmt.Println(i4)

	// 型の違う値を代入することはできない
	// i4 = "Hello"
	// fmt.Println(i4)

	outer()

}
