package main

import (
	"fmt"
)

func main(){

	// // 無限ループ
	// for{
	// 	fmt.Println("無限ループ")
	// }

	// forループ
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 3{
			continue
		}
		if i == 8{
			break
		}
	}

	// forループ(配列)
	arr := [3]int{1,2,3}
	for i:=0;i<len(arr);i++ {
		fmt.Println("forループ(配列)", i)
	}

	// foreach(配列)
	for i, v := range arr{
		fmt.Printf("foreach index = %d, value = %d\n",i, v)
	}

	// foreach(map型)
	m := map[string] int{"apple":100, "banana":200}
	for k, v := range m{
		fmt.Println(k,v)
	}

	// ラベル付for
	Loop:
		for{
			for{
				for{
					fmt.Println("START")
					break Loop		// Loopの部分まで処理を飛ばす
				}
				fmt.Println("処理をしない")
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("END")
	
}