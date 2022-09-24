package main

import (
	"bufio"
	"os"
	"regexp"
	"time"
)

// VirtualIO - 入出力を模擬したもの
type VirtualIO struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer

	inputFilePath string
	inputLines    []string
	pollingTime   time.Duration
}

// 新規作成
//
// - 行読取
//
// Parameters
// ----------
// setVIO - 初期化に使える
func NewVirtualIO() *VirtualIO {
	// 実体をメモリ上に占有させる
	//
	// - 規定値：標準入出力
	var virtualIo = VirtualIO{
		scanner:       bufio.NewScanner(os.Stdin),
		writer:        bufio.NewWriter(os.Stdout),
		inputFilePath: "",
		inputLines:    []string{},
		// 1秒は長いが、しかたない
		pollingTime: 1 * time.Second,
	}

	// virtualIo.Scanner.Split(bufio.ScanWords) // 空白で区切る
	virtualIo.scanner.Split(bufio.ScanLines) // 改行で区切る
	// 入力バッファーのサイズを巨大にする
	virtualIo.scanner.Buffer([]byte{}, 100000007)

	// バーチャルIOのアドレスを返す
	return &virtualIo
}

// IsEmpty - 空っぽか
func (vio *VirtualIO) IsEmpty() bool {
	// １行以上存在し、０行目が空行なら、詰める
	for len(vio.inputLines) != 0 && vio.inputLines[0] == "" {
		vio.inputLines = vio.inputLines[1:len(vio.inputLines)]
	}

	// ０行なら空っぽ
	return len(vio.inputLines) == 0
}

// ReplaceInputToFileLines - 標準入力を使うのをやめ、ファイルの先頭行から１行ずつ切り取る方法に変えます
//
// Parameters
// ----------
// inputFilePath - ファイルパス
func (vio *VirtualIO) ReplaceInputToFileLines(inputFilePath string) {
	vio.inputFilePath = inputFilePath
}

var re = regexp.MustCompile("\r\n|\n")

func (vio *VirtualIO) ScannerScan() bool {

	// テキストファイルから読み込むなら
	if vio.inputFilePath != "" {

		var popAllLines = func() []string {
			// ファイル読込
			var bytes, err = os.ReadFile(vio.inputFilePath)
			if err != nil {
				panic(err)
			}

			var text = string(bytes)

			// ファイルを空にする
			os.Truncate(vio.inputFilePath, 0)

			// 全文を改行でスプリット
			return re.Split(text, -1)
		}

		// バッファーが空なら、ファイルから取ってくる
		if vio.IsEmpty() {
			// 全行取得
			vio.inputLines = popAllLines()
		}

		// バッファーが空の間ブロック（繰り返し）する
		for vio.IsEmpty() {
			// スリープする。なぜなら、入力がないときブロックするという機能を入れないと、無限に空文字列を読み続けてしまうから
			time.Sleep(vio.pollingTime)

			// 全行取得
			vio.inputLines = popAllLines()
		}

		return true
	}

	return vio.scanner.Scan()
}

func (vio *VirtualIO) ScannerText() string {

	// テキストファイルから読み込むなら
	if vio.inputFilePath != "" {
		// 先頭の１行を取り出し
		var firstLine = vio.inputLines[0]

		// 繰り上がり
		vio.inputLines = vio.inputLines[1:len(vio.inputLines)]

		return firstLine
	}

	return vio.scanner.Text()
}

func (vio *VirtualIO) WriterFlush() {
	virtualIo.writer.Flush()
}
