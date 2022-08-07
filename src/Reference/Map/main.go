package main

import (
	"fmt"
)

func main() {
	var slice []int
	fmt.Println(slice)

	var slice2 []int = []int{100, 200}
	fmt.Println(slice2)

	slice2[0] = 1000
	fmt.Println(slice2)

	slice3 := []string{"A", "B"}
	fmt.Println((slice3))

	slice4 := make([]int, 5)
	fmt.Println(slice4) // 0 0 0 0 0

	slice5 := []int{0, 1, 2, 3, 4}
	// 要素1~2の取り出し(先頭が以上，末尾が未満になっている！)
	fmt.Println(slice5[1:3]) // 1 2
	// 要素3より小さいものの取り出し
	fmt.Println(slice5[:3]) // 0 1 2
	// 要素3以上の取り出し
	fmt.Println(slice5[3:]) // 3 4
	// 最後の要素の取り出し
	fmt.Println(slice5[len(slice5)-1]) // 4

	// append
	sl := []int{100, 200}
	sl = append(sl, 300)
	fmt.Println(sl) // 100 200 300

	// append(複数)
	sl = append(sl, 400, 500)
	fmt.Println(sl) // 100 200 300 400 500

	// スライスの代入
	sl2 := sl
	sl2[0] = 1000    // 参照型でコピーされるので代入されると元の配列も影響を受ける
	fmt.Println(sl)  // 1000 200 300 400 500
	fmt.Println(sl2) // 1000 200 300 400 500

	// copy
	copyed := make([]int, 5) // 先にコピー先をmakeで作成
	copyedNum := copy(copyed, sl)
	fmt.Println("コピーされた数", copyedNum)

	copyed[0] = 10000
	fmt.Println("src", sl)        // 1000 200 300 400 500		←ソースには影響していない．
	fmt.Println("copyed", copyed) // 10000 200 300 400 500

	// froeach
	for index, value := range sl {
		fmt.Println("index=", index, "value=", value)
	}

	// Sum関数の呼び出し *スライスを引数にする場合「...」を変数の後に置く
	fmt.Println(Sum(sl...))
}
