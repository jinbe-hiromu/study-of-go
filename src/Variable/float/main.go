package main

import (
	"fmt"
)

func main(){
	// float型は実行環境に依存せずfloat64となる
	fl := 1.2
	fmt.Println(fl)
	fmt.Printf("%T\n", fl)		// float64
}