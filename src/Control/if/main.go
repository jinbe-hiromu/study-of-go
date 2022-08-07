package main

import (
	"fmt"
)

func main(){
	a := 1
	if a == 1{
		fmt.Println("true")
	}	else		// 必ずワンラインで } elseと記述が必要。改行はNG
	{
		fmt.Println("false")
	}
}