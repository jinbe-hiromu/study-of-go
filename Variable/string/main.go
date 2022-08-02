package main

import (
	"fmt"
)

func main(){
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Println(s[0])		// バイト型が出力(72)
	fmt.Println(string(s[0]))		// １文字が出力
	
	// バッククオートで改行文字列を出力できる
	fmt.Println(`
	test
	test
	test`)
}