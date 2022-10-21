
package test

import (
	"testing"

	"github.com/oyamo/telegram-compiler/src"
)

const (
	kotlin_test_code = `fun main() {
		val kotlin = "🙂"
		println(kotlin)
	 }`
)

func TestKotlin(t *testing.T) {
	res, err := src.Kotlin(kotlin_test_code)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("Empty response")
	}

	if res.StatusCode != 200 {
		t.Error("Invalid status code")
	}

	t.Log(res)
}
