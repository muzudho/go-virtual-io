// BOF [O1o4o0]

package main

import "fmt"

// 文字列出力
func (vio *VirtualIO) Printf(format string, a ...interface{}) {
	fmt.Fprintf(vio.writer, format, a...)
}

// EOF [O1o4o0]
