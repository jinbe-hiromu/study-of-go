package main

import (
	"fmt"
)

// panic：goの処理を強制的に終了

func main(){
	defer func(){
		// 復帰処理。panic後に実行される処理
		if x:= recover(); x!=nil{
			fmt.Println("panic occured")		// こちらの文字列が出力される
		}
	}()
	panic("panic!!!")
	fmt.Println("Start")

}