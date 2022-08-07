package main

import (
	"fmt"
)

func init(){
	fmt.Println("init")		// mainよりも先に出力される  init()は特別な関数名。初期化処理に使われる
}

func main(){
	fmt.Println("main")
}