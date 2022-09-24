package main

import (
	"bufio"
	"os"
	"time"
)

// VirtualIO - 入出力を模擬したもの
type VirtualIO struct {
	scanner       *bufio.Scanner
	writer        *bufio.Writer
	inputFilePath string
	inputText     string
}

// 新規作成
//
// - 行読取
func NewVirtualIO() *VirtualIO {
	// 実体をメモリ上に占有させる
	//
	// - 規定値：標準入出力
	var virtualIo = VirtualIO{
		scanner:       bufio.NewScanner(os.Stdin),
		writer:        bufio.NewWriter(os.Stdout),
		inputFilePath: "",
		inputText:     "",
	}

	// virtualIo.Scanner.Split(bufio.ScanWords) // 空白で区切る
	virtualIo.scanner.Split(bufio.ScanLines) // 改行で区切る
	// 入力バッファーのサイズを巨大にする
	virtualIo.scanner.Buffer([]byte{}, 100000007)

	// バーチャルIOのアドレスを返す
	return &virtualIo
}

// Stubs Stdin in 'fn'
// See also: 📖 [golangのAtCoder向けデバック方法(VSCode)](https://qiita.com/tasmas/items/d2d5a8c95fa48e415702)
//
// Examples
// --------
// inbuf := "入力されたつもりの文字列。テキストファイルから読み込んでくる"
//
//	StubStdin("ファイル名", func() {
//	    main()
//	})
//
// Parameters
// ----------
// textToWrite - 書き込みたい文字列
func (vio *VirtualIO) SetupStubStdin(inputFilePath string) {
	vio.inputFilePath = inputFilePath
}

func (vio *VirtualIO) ScannerScan() bool {

	// テキストファイルから読み込むなら
	if vio.inputFilePath != "" {

		var getText = func() string {
			// ファイル読込
			var bytes, err = os.ReadFile(vio.inputFilePath)
			if err != nil {
				panic(err)
			}

			return string(bytes)
		}

		// 文字列取得
		vio.inputText = getText()

		// 空文字でなくなるまでブロック（繰り返し）する
		for vio.inputText == "" {
			// TODO 入力がないときブロックするという機能を入れないと、無限に空文字列を読み続けてしまう。1秒は長いが、しかたない
			time.Sleep(1 * time.Second)

			// 文字列取得
			vio.inputText = getText()
		}

		return true
	}

	return vio.scanner.Scan()
}

func (vio *VirtualIO) ScannerText() string {

	// テキストファイルから読み込むなら
	if vio.inputFilePath != "" {
		return vio.inputText
	}

	return vio.scanner.Text()
}

func (vio *VirtualIO) WriterFlush() {
	virtualIo.writer.Flush()
}
