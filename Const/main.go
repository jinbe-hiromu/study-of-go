package main

import (
	"fmt"
)

// 定数宣言(型推論)
const Pi = 3.14

// 連続して定数を宣言する場合
const(
	URL = "http://xxx.co.jp"
	SiteName = "test"
)

const(
	c0 = iota			// 1 連番を付与 
	c1						// 2
	c2						// 3
)

func main(){
	fmt.Println(Pi)
	fmt.Printf("%T\n",Pi)		// float64

	fmt.Println(SiteName)
	fmt.Printf("%T\n",SiteName)		// string

	fmt.Println(URL)
	fmt.Printf("%T\n",URL)		// string

	fmt.Println(c0,c1,c2)		// 0 1 2
	fmt.Printf("%T\n",c0)		// int
}