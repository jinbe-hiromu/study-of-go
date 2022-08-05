package main

import (
	"strconv"
	"fmt"
)

func main(){
	s := "100"

	i, err := strconv.Atoi(s)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%T\n",i)
}