package main

import (
	"os"
	"fmt"
)

func TestDefer1(){
	defer fmt.Println("END")		// deferで登録したものは最後に実行される
	fmt.Println("SATART")

	// START
	// END
}

func TestDefer2(){
	// 最後にdefer登録されたものから実行される。
	defer fmt.Println("END1")		// 3番目に実行
	defer fmt.Println("END2")		// 2番目に実行
	defer fmt.Println("END3")		// 1番目に実行

	fmt.Println("SATART")

	// START
	// END3
	// END2
	// END1
}

func TestDefer3(){
	// 無名関数で記述すると、順番通りに実行される。
	defer func() {
		fmt.Println("END1")		// 1番目に実行
		fmt.Println("END2")		// 2番目に実行
		fmt.Println("END3")		// 3番目に実行
	}()
	fmt.Println("SATART")

	// START
	// END1
	// END2
	// END3
}

func CreateText(){
	file,err := os.Create("test.txt")
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()		// ファイルのDispose処理で使用する

	file.Write([]byte("Hello"))
}

func main(){
	TestDefer1()
	TestDefer2()
	TestDefer3()
	CreateText()
}