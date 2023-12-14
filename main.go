package main

import "fmt"

// 型
// 型変換

func main() {
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)               // 1
	// fmt.Printf("i = %T\n", i)       // i = int
	// fmt.Printf("fl64 = %T\n", fl64) // fl64 = float64

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2) // i2 = int

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3 = %T\n", i3) // i3 = int
	// fmt.Println(i3)             // 10

	// string -> int
	// var s string = "A"
	// fmt.Printf("s = %T\n", s)

	// i, err := strconv.Atoi(s) // 文字列型からint型に変換
	// if err != nil {
	// 	fmt.Println(err) // strconv.Atoi: parsing "A": invalid syntax
	// }

	// fmt.Println(i)
	// fmt.Printf("i = %T\n", i) // i = int

	// var i2 int = 200
	// fmt.Printf("i2 = %T\n", i2) // i2 = int

	// s2 := strconv.Itoa(i2)      // int型から文字列型に変換
	// fmt.Printf("s2 = %T\n", s2) // s2 = string

	var h string = "Hello World"
	b := []byte(h) // 文字列をバイト型に変換
	fmt.Println(b) // [72 101 108 108 111 32 87 111 114 108 100]

	h2 := string(b) // バイト型を文字列に変換
	fmt.Println(h2) // Hello World
}
