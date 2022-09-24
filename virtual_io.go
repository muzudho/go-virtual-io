package main

import (
	"bufio"
	"os"
)

// VirtualIO - 入出力を模擬したもの
type VirtualIO struct {
	Scanner *bufio.Scanner
	Writer  *bufio.Writer
}

// 新規作成
//
// - 行読取
func NewVirtualIO() *VirtualIO {
	// 実体をメモリ上に占有させる
	var virtualIo = VirtualIO{
		Scanner: bufio.NewScanner(os.Stdin),
		Writer:  bufio.NewWriter(os.Stdout),
	}

	// virtualIo.Scanner.Split(bufio.ScanWords) // 空白で区切る
	virtualIo.Scanner.Split(bufio.ScanLines) // 改行で区切る
	// 入力バッファーのサイズを巨大にする
	virtualIo.Scanner.Buffer([]byte{}, 100000007)

	// バーチャルIOのアドレスを返す
	return &virtualIo
}
