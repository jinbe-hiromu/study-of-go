package main

import (
	"fmt"
)

func main(){
	// interface型・・・Variant型のようなもの。全ての型を代入できる
	// 初期値はnil
	// ただし、演算はできない
	var x interface{}
	fmt.Println(x)		// <nil>

	x = 1
	fmt.Println(x)		// 1
	// fmt.Println(x+1)		// NG

	x = "string"
	fmt.Println(x)		// string
}