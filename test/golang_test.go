package test

import (
	"testing"

	"github.com/oyamo/telegram-compiler/src"
)

const (
	test_code = `package main
	import "fmt"
	func main() {
		fmt.Println("Hello, 世界")
	}`
)

func TestGolang(t *testing.T) {
	res, err := src.Golang(test_code)
	if err != nil {
		t.Error(err)
	}

	if res == "" {
		t.Error("Empty response")
	}

	t.Log(res)
}
