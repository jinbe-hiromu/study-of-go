package main

import (
	"fmt"
)

// switch型スイッチ
func ExecuteBranch(x interface{}){
	// 引き数の型によって処理を記述可能
	switch v := x.(type) {
	case int:
		fmt.Println("Type is int. value = ",v)
	case string:
		fmt.Println("Type is string. value = ",v)
		default:
		fmt.Println("I don't know")
	}
}

func main(){

	n:=1
	switch n{
	case 1,2:
		fmt.Println("1 or 2")
	case 3,4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("default")
	}

	// caseのあとに評価式を取るパターン
	// 評価式と固定値は混ぜてcaseに記述することはできない。
	switch {		// 固定値の場合はswitch nと記述していた点に注意！
	case 1 <= n && n < 2:
		fmt.Println("1 or 2")
	case 3 <= n && n <= 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("default")
	}

	ExecuteBranch(3)
	ExecuteBranch("str")
}