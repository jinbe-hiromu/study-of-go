package main

import (
	"strconv"
	"fmt"
)

func main(){
	// string → int変換
	str1 := "100"
	convertedInt, _ := strconv.Atoi(str1)		// 2つ目の返り値はError。_(破棄)すれば使用しなくてもコンパイルエラーが発生しない
	fmt.Println(convertedInt)
	fmt.Printf("%T\n",convertedInt)		// int

  // int → string変換
	int2 := 100
	contertedStr := strconv.Itoa(int2)
	fmt.Println(contertedStr)
	fmt.Printf("%T\n",contertedStr)		// string

	contertedStr2 := string(int2)				// バイト配列→string
	fmt.Println(contertedStr2)					// 「d」(ASCIIの100)
	fmt.Printf("%T\n",contertedStr2)		// string

	// byte配列 → string
	str3 := "Hello world"
	byteArray := []byte(str3)
	fmt.Println(byteArray)					// [72 101 108 108 111 32 119 111 114 108 100]
	fmt.Printf("%T\n",byteArray)		// []uint8

	str4 := string(byteArray)
	fmt.Println(str4)								// Hello world
	fmt.Printf("%T\n",str4)					// string

}