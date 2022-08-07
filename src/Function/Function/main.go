package main

import (
	"fmt"
)

func Plus(x int, y int) int{
	return x + y
}

// 引数が同じ型の場合省略可能
// 返り値が２つの場合()でくくる
func Div(x,y int) (int, int){
	q := x / y
	r := x % y
	return q,r
}

// 返り値の変数を定義するとreturnで返却可能
func Double(price int) (result int){
	result = price * 2
	return
}

func main(){
	i := Plus(1,2)
	fmt.Println(i)

	sho, amari := Div(9,4)
	fmt.Println(sho, amari)

	ret := Double(1000)
	fmt.Println(ret)
}