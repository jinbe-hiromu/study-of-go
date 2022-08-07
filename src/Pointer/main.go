package main

import (
	"fmt"
)

// C,C++と同じように＊でポインタ型が使用できる．
func Double(i *int) {
	*i *= 2
}

func main() {
	i := 100
	Double(&i)
	fmt.Println(i) //200
}
