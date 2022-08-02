package main

import (
	"fmt"
)

func main(){
	var arr [3]int
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Printf("%T\n",arr)

	var strArray [3]string = [3]string{"A","B"}
	fmt.Println(strArray)
	fmt.Printf("%T\n",strArray)

	var intArray [3]int = [3]int{1,2,3}
	fmt.Println(intArray)
	fmt.Printf("%T\n",intArray)

	// 要素数を省略する方法
	strArrayShoryaku := [...]string{"A","B","C"}
	fmt.Println(strArrayShoryaku)
	fmt.Printf("%T\n",strArrayShoryaku)

	// 要素数を調べる方法
	fmt.Println(len(strArrayShoryaku))
}