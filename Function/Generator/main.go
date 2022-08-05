package main

import (
	"fmt"
)

// ジェネレータ：クロージャ―を応用して連続した値を返却する
func GetGenerator() func() int{
	counter := 0
	return func() int{
		counter++
		return counter
	}
}

func main(){
	generator := GetGenerator()
	fmt.Println(generator())		// 1
	fmt.Println(generator())		// 2
	fmt.Println(generator())		// 3
	fmt.Println(generator())		// 4
	fmt.Println(generator())		// 5

	generator1 := GetGenerator()
	fmt.Println(generator1())		// 1
	fmt.Println(generator1())		// 2
	fmt.Println(generator1())		// 3
	fmt.Println(generator1())		// 4
	fmt.Println(generator1())		// 5
}