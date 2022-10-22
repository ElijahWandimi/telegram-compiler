package test

import (
	"fmt"
	"testing"

	"github.com/oyamo/telegram-compiler/src"
)

const (
	clang_test_code = `#include <stdio.h>
	int main() {
		printf("Hello, World!");
		return 0;
		}`
)

func TestClang(t *testing.T) {
	res, err := src.Clang(java_test_code)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("Empty response")
	}

	if res.StatusCode != 200 {
		t.Error("Invalid status code " +fmt.Sprint(res.StatusCode))
	}

	t.Log(res.Body)
}
