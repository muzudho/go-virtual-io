package main

import (
	"strconv"
)

// グローバル変数として、バーチャルIOを１つ新規作成
// アプリケーションの中では 標準入出力は これを使うようにする
var virtualIo = NewVirtualIO()

func main() {
	// この関数を抜けるときに、バーチャルIOの出力バッファーをフラッシュする
	defer virtualIo.WriterFlush()

	// 入力を読取る
	if virtualIo.ScannerScan() {
		var text = virtualIo.ScannerText()
		var i, err = strconv.Atoi(text)
		if err != nil {
			panic(err)
		}

		virtualIo.Printf("%d is ok\n", i) // 出力
	}
}
