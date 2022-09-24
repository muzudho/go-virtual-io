# go-virtual-io

delveが標準入力でフリーズするので、デバッグ中は、標準入力をファイル入力へ置き換える仕組みがほしい  

そこで、  
テストでは 標準入力をファイル入力へ置き換えるようにし、  
テストをデバッグ実行するようにする  

# Test

## データファイル編集 - test.input.txt

👇 以下の既存ファイルを編集してほしい

```plaintext
    📂
    └── 📄 test.input.txt
```

```plaintext
10
```

`*.input.txt` というファイル名は、内容が読み取られるとともに空っぽに消される目印にしている。注意してほしい  

## テスト実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```plaintext
go test
```

Output:  

```plaintext
10 is ok
PASS
ok      github.com/muzudho/go-virtual-io        0.194s
```

## テストのデバッグ実行

📄 `test.input.txt` を再度編集し、 `main_test.go` をデバッグ実行してほしい  

# Run

## 実行

👇 以下のコマンドをコピーして、ターミナルに貼り付けてほしい  

Input:  

```plaintext
go run .
11
```

Output:  

```plaintext
11 is ok
```

## 参考にした記事

📖 [golangのAtCoder向けデバック方法(VSCode)](https://qiita.com/tasmas/items/d2d5a8c95fa48e415702)  
📖 [scanner.Scan() hangs in GoLand debugger](https://stackoverflow.com/questions/53461228/scanner-scan-hangs-in-goland-debugger)  

### 文字列操作

📖 [Go言語 Split 文字列を分割して配列にする](https://itsakura.com/golang-split)  

### ファイル操作

📖 [How to Truncate a File in Golang?](https://www.geeksforgeeks.org/how-to-truncate-a-file-in-golang/)  

EOF