# Step [O1o0] How to make

## Step [O1o1o0] git向け対応 - .gitignore ファイル

👇 以下のファイルが既存なら編集してほしい。無ければ新規作成してほしい  

```plaintext
  	📂 kifuwarabe-uec14
👉  └── 📄 .gitignore
```

👇 冒頭に追加してほしい  

```plaintext
# この下に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...

# [O1o1o0]
*.input.txt

# この上に kifuwarabe-uec14 でリポジトリにコミットしないものを追加する
# ---------------------------------------------------------------
# ...略...
```

## Step [O1o2o0] データファイル編集 - test.input.txt

👇 以下のファイルを新規作成してほしい  

```plaintext
    📂
    ├── 📄 .gitignore
👉  └── 📄 test.input.txt
```

```plaintext
10
```

`*.input.txt` というファイル名は、内容が読み取られるとともに空っぽに消される目印にしている。消えて困る内容を書かないように注意してほしい  

## Step [O1o3o0] バーチャルIO作成 - virtual_io.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂
    ├── 📄 .gitignore
  	├── 📄 test.input.txt
👉  └── 📄 virtual_io.go
```

```go
// BOF [O1o3o0]

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

// EOF [O1o3o0]
```

## Step [O1o4o0] バーチャルIO作成 - virtual_io_fmt.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂
    ├── 📄 .gitignore
  	├── 📄 test.input.txt
👉 	├── 📄 virtual_io_fmt.go
    └── 📄 virtual_io.go
```

```go
// BOF [O1o4o0]

package main

import "fmt"

// 文字列出力
func (vio *VirtualIO) Printf(format string, a ...interface{}) {
	fmt.Fprintf(vio.writer, format, a...)
}

// EOF [O1o4o0]
```

## Step [O1o5o0] ファイル作成 - main.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂
    ├── 📄 .gitignore
👉  ├── 📄 main.go
  	├── 📄 test.input.txt
 	├── 📄 virtual_io_fmt.go
    └── 📄 virtual_io.go
```

```go
// BOF [O1o5o0]

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

// BOF [O1o5o0]
```

## Step [O1o6o0] ファイル作成 - main_test.go ファイル

👇 以下のファイルを新規作成してほしい  

```plaintext
  	📂
    ├── 📄 .gitignore
👉  ├── 📄 main_test.go
    ├── 📄 main.go
  	├── 📄 test.input.txt
 	├── 📄 virtual_io_fmt.go
    └── 📄 virtual_io.go
```

```go
// BOF [O1o6o0]

package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	virtualIo.ReplaceInputToFileLines("./test.input.txt")
	main()
}

// EOF [O1o6o0]
```

## Step [O1o7o0] モジュール作成

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go mod init github.com/muzudho/go-virtual-io
#           --------------------------------
#           1
# 1. モジュール名。この部分はあなたのレポジトリに合わせて変えてほしい
```

👇 以下のファイルが自動生成される  

```plaintext
  	📂
    ├── 📄 .gitignore
👉  ├── 📄 go.mod
    ├── 📄 main_test.go
    ├── 📄 main.go
  	├── 📄 test.input.txt
 	├── 📄 virtual_io_fmt.go
    └── 📄 virtual_io.go
```

```plaintext
module github.com/muzudho/go-virtual-io

go 1.19
```

## Step [O1o8o0] もしワークスペースズモードを使っているなら

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```shell
go work use .
go mod tidy
```
