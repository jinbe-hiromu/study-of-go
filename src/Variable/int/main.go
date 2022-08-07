package main

import (
	"fmt"
)

func main(){
	// int型は実行環境によって異なる
	var i int = 100
	
	// uint8	8ビット符号なし整数	0 ～ 255
	// uint16	16ビット符号なし整数	0 ～ 65535
	// uint32	32ビット符号なし整数	0 ～ 4294967295
	// uint64	64ビット符号なし整数	0 ～ 18446744073709551615

	// 型の出力
	fmt.Printf("%T\n",i)

	// キャスト
	fmt.Printf("%T\n",int32(i))
}