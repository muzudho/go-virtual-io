package main

import (
	"bufio"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	inbuf := readFile("./test.txt")
	stubStdin(inbuf, func() {
		main()
	})
}

func readFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Stubs Stdin in 'fn'
// See also: 📖 [golangのAtCoder向けデバック方法(VSCode)](https://qiita.com/tasmas/items/d2d5a8c95fa48e415702)
//
// Examples
// --------
// inbuf := "入力されたつもりの文字列。テキストファイルから読み込んでくる"
//
//	stubStdin(inbuf, func() {
//	    main()
//	})
//
// Parameters
// ----------
// textToWrite - 書き込みたい文字列
func stubStdin(inputText string, fn func()) {
	// これより、ラムダ計算の専門用語で η（イータ）簡約 と呼ばれることと同じ考え方を利用する。
	// Input ストリームと使い勝手が同等になるよう、 Read モードと Write モードのファイル（メモリ上に存在する）を取得
	inr, inw, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	// Input ストリームに書き込んでいるつもりで、 Write モードのファイルに書き込む
	_, _ = inw.Write([]byte(inputText))
	// 書込みをフラッシュして終わる
	inw.Close()

	// Input ストリームから読込んでいるつもりで、 Read モードのファイルを `os.Stdin` と差し替える
	os.Stdin = inr
	// このスキャナーは、標準入力をスキャンしているように見えて、メモリ上に存在するファイルをスキャンしている
	virtualIo.Scanner = bufio.NewScanner(os.Stdin)

	// あとは ふつうに処理を行う
	fn()

	// TODO `os.Stdin` を元に戻さなくていいのか？ fn() が main() プログラムと同等で、あとは終了するるだけなら 良いとはいえるが
}
