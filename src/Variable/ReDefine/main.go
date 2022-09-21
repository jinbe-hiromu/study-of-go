package main

import (
	"fmt"
)

func main(){
	s := "Hello Golang"

	if s:= true; s {
		fmt.Println(s)			// true
	}

	fmt.Println(s)				// Hello Golang
}

