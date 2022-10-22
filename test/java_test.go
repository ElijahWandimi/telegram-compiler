package test

import (
	"fmt"
	"testing"

	"github.com/oyamo/telegram-compiler/src"
)

const (
	java_test_code = `public class Main {
		public static void main(String[] args) {
			System.out.println("Hello, 世界");
			}
		}`
)

func TestJava(t *testing.T) {
	res, err := src.Java(java_test_code)
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
