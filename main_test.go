package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Hello, playground")
	os.Exit(m.Run())
}
