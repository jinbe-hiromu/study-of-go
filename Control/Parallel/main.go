package main

import (
	"time"
	"fmt"
)

func Sub(){
	i:=0
	for{
		fmt.Println("Sub thread",i)
		time.Sleep(1000 * time.Millisecond)
		i++
	}
}

func main(){
	go Sub()		// 平行処理は関数の前にgoを付けるだけ

	i := 0
	for{
		fmt.Println("Main thread",i)
		time.Sleep(2000 * time.Millisecond)
		i++
	}

}