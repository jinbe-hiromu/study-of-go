package main

import (
	"fmt"
	"time"
)

func main() {
	// チャネル：goルーチン間でデータのやり取りするキュー
	// chanになにも<-がかかれていない場合，送受信可能という意味
	channel1 := make(chan int, 5) // 第二引数はキューのサイズを指定．何もないバッファが0でキューに詰めるとデッドロックエラーが出力される．
	channel1 <- 1
	fmt.Println("len=", len(channel1)) // len=1
	fmt.Println(<-channel1)            // 取り出し
	fmt.Println("len=", len(channel1)) // len=0		取り出し後は数が-1される
	// fmt.Println(<-channel1)            // 取り出し（キューが0個の状態で取り出しをするとエラー

	// 受信専用チャネル
	// channel2 := make(<-chan int)

	// 送信専用チャネル
	// channel3 := make(chan<- int)

	// string型のチャンネル
	channelStr := make(chan string, 5)
	channelStr <- "ABC"
	channelStr <- "DEF"
	fmt.Println(<-channelStr) // ABC
	fmt.Println(<-channelStr) // DEF

	// time.MIlliseconedの値
	fmt.Println("time.Millisecond=", time.Millisecond) // 1ms←単位まで出力される

	// チャネルとゴルーチン
	testChannel := make(chan int) // これはバッファを宣言しなくてもエラーにならなかった．なぜ？　行12ではバッファを書かないとエラーになっていたが．．．main関数のみで使うときはバッファサイズが必要なのか？
	go Receive(testChannel)

	// 送信
	for i := 0; i < 100; i++ {
		testChannel <- i
		fmt.Println("送信：", i)
		// time.Sleep(xxx) // xxxはnsで指定

		time.Sleep(10 * time.Millisecond) // 10ms
	}
	// channelのクローズ処理
	close(testChannel)
	time.Sleep(3 * time.Second) // ゴルーチンが修了するまで3秒だけ待機

	// Selectを使用した受信関数
	testChannel1 := make(chan int)
	testChannel2 := make(chan int)

	go ReceiveSelect(testChannel1, testChannel2)

	for i := 0; i < 100; i++ {
		fmt.Println("送信：", i)
		testChannel2 <- i
		testChannel1 <- i
		time.Sleep(10 * time.Millisecond)
	}
}

// 受信関数
func Receive(channel chan int) {
	for {
		value, isAvairable := <-channel
		if !isAvairable {
			fmt.Println("END")
			break
		}
		fmt.Println("受信：", value) // channelから受信した数値を出力
	}
}

// Selectを使用したReceive関数
func ReceiveSelect(channel1 chan int, channel2 chan int) {
	for {
		// 先に受信したcaseの方を実行する
		select {
		case i1 := <-channel1:
			fmt.Println("receive channel1:", i1)
		case i2 := <-channel2:
			fmt.Println("receive channel2:", i2)
		}
	}
}
