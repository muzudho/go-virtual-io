package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	StubStdin("./test.txt", func() {
		main()
	})
}
