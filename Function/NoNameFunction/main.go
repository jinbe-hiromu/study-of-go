package main

import (
	"fmt"
)

func main(){
	// ラムダ式
	f:= func(x, y int) int{
		return x + y
	}

	i := f(1,2)
	fmt.Println(i)

	// 事前に引数を渡すやり方
	i2 := func(x, y int) int{
		return x + y
	}(1,2)

	fmt.Println(i2)
}