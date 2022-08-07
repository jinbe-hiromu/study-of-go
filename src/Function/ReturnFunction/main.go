package main

import (
	"fmt"
)

// 返り値なし
func ReturnFunc() func(){
	return func(){
		fmt.Println("I'm a function1")
	}
}

// 返り値あり
func ReturnFuncWithReturnVal() func() int{
	return func() int{
		fmt.Println("I'm a function2")
		return 2
	}
}

// 引数あり
func ReturnFuncWithArg() func(x int){
	return func(x int){
		fmt.Println("I'm a function3")
		fmt.Println(x*2)
	}
}


func main(){
	function := ReturnFunc()
	function()

	function1 := ReturnFuncWithReturnVal()
	val := function1()
	fmt.Println(val)

	function3 := ReturnFuncWithArg()
	function3(3)
}