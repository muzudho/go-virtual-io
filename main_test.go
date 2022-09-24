package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	inbuf := readFile("./test.txt")
	StubStdin(inbuf, func() {
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
