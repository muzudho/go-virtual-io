package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	virtualIo.SetupStubStdin("./test.txt")
	main()
}
