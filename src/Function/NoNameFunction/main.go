package main

import (
	"fmt"
)

func main(){
	// 何回も呼び出す場合
	f := func(x,y int) int{
		return x + y
	}
	i := f(1,2)		// x,yを渡している
	fmt.Println(i)

	// 1回のみ呼び出す場合
	ret := func(i,y int) int{
		return i *y
	}(1,10)			// ここで引数を渡す
	fmt.Println(ret)
}