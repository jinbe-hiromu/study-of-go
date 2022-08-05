package main

import (
	"fmt"
)

func Later() func(string) string{
	var store string			// 保持するための変数
	return func(next string) string{
		s:=store			// 保持されているものをreturnする
		store=next		// 引数でわたってきたものを保持
		return s
	}
}

func main(){
	f := Later()								// fが存在する間、storeにいれた文字列は削除されない。
	fmt.Println(f("Hello"))			// 空文字（最初に保持されているのが空文字のため）
	fmt.Println(f("My"))				// Hello
	fmt.Println(f("name"))			// My
	fmt.Println(f("is"))				// name
	fmt.Println(f("Kotone"))		// is
}