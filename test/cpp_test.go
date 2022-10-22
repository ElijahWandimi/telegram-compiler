package test

import (
	"fmt"
	"testing"

	"github.com/oyamo/telegram-compiler/src"
)

const (
	cpp_test_code = `#include <iostream>
	using namespace std;
	int main() {
		cout << "Hello, World!";
		return 0;
	}`
)

func TestCplus(t *testing.T) {
	res, err := src.CPlus(cpp_test_code)
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
