package main

import (
	"fmt"
)

func main() {
	// 明示的
	var map1 = map[string]int{"A": 100, "B": 200}
	fmt.Println(map1)

	// 暗黙的
	map2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(map2)

	// makeでmapを生成
	map3 := make(map[string]int)

	// 値の追加
	map3["A"] = 100
	map3["B"] = 200
	fmt.Println(map3)

	// 値の変更
	map3["A"] = 1000
	fmt.Println(map3)

	// 値の削除
	delete(map3, "A")
	fmt.Println(map3)

	// キーがない場合の値
	fmt.Println(map3["C"]) // 空文字　mapで指定したキーが存在しない場合，空文字が出力される

	// 値取得に失敗した場合のエラーハンドリング
	value, ret := map3["C"]
	if !ret {
		fmt.Println("error") // キーが存在しないので，こちらが出力
	} else {
		fmt.Println(value)
	}

	// foreach
	map4 := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	for k, v := range map4 {
		fmt.Println(k, v)
	}
}
