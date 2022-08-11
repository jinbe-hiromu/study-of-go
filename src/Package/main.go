package main

import (
	"fmt"
	"main/foo"		// go.modがあるのでモジュールモード。モジュールモードで内部パッケージを参照する場合、モジュール名(main)＋パッケージ名(foo)と記述する
)

func main() {
	fmt.Println(foo.Max)
}