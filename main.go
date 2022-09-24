package main

import (
	"fmt"
	"strconv"
)

// グローバル変数として、バーチャルIOを１つ新規作成
// アプリケーションの中では 標準入出力は これを使うようにする
var virtualIo = NewVirtualIO()

func main() {
	// この関数を抜けるときに、バーチャルIOの出力バッファーをフラッシュする
	defer virtualIo.Writer.Flush()

	// 入力を読取る
	if virtualIo.Scanner.Scan() {
		var text = virtualIo.Scanner.Text()
		var i, err = strconv.Atoi(text)
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(virtualIo.Writer, fmt.Sprintf("%d is ok", i)) // 出力
	}
}
