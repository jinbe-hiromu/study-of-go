package main

import (
	"fmt"
)

func ExcecuteFunction(f func()){
	f()
}

func main(){
	ExcecuteFunction(func(){
		fmt.Println("引数で渡したFunctionです")
	})
}