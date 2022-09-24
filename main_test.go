package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	virtualIo.ReplaceInputToFileLines("./test.input.txt")
	main()
}
