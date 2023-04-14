package test

import (
	"fmt"
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
		t.Log(res.Body)
	} else {
		t.Log(res.Body)

		if res.StatusCode != 200 {
			t.Error("Invalid status code " +fmt.Sprint(res.StatusCode))
		}
	
	}
}
