package test

import (
	"fmt"
	"testing"

	"github.com/oyamo/telegram-compiler/src"
)

const (
	python_test_code = `
	print("Hello, 世界")`
)

func TestPython(t *testing.T) {
	res, err := src.Python(python_test_code)
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
