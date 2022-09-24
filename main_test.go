package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	virtualIo.SetInputFromFile("./test.txt")
	main()
}
