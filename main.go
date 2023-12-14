package main

import "fmt"

// 型
// 配列型

/**
** Goでは配列型の要素数を変更できない
** 要素の追加を行う際はsliceを使う
**
** 配列型 = [要素数]型 要素を変更できない
** スライス型 = []型 要素を変更できる
**/

func main() {
	var arr1 [3]int
	fmt.Println(arr1)        // [0 0 0]
	fmt.Printf("%T\n", arr1) // [3]int

	// sliceを使わないとGoでは配列の長さを変更できない

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2) // [A B ]

	var arr3 = [3]int{1, 2, 3}
	fmt.Println(arr3) // [1 2 3]

	var arr4 = [...]int{4, 5}
	fmt.Printf("%T\n", arr4) // [2]int

	// 配列の一部を取り出す
	fmt.Println(arr2[0]) // A
	fmt.Println(arr2[1]) // B

	// 配列の最後の要素を取り出す
	fmt.Println(arr2[len(arr2)-1]) // B

	// 配列の要素を変更する
	arr2[2] = "C"
	fmt.Println(arr2) // [A B C]

	// var arr5 [4]int
	// arr5 = arr1       // 配列の長さが同じでないため代入できない
	// fmt.Println(arr5) // [0 0 0]
}
