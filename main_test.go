package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	virtualIo.SetInputFromFile("./test.input.txt")
	main()
}
