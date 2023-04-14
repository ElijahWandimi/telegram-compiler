package main

import (
	"log"
	"github.com/oyamo/telegram-compiler/src"
)
func main() {
	err := src.CreateServer()
	if err != nil {
		log.Fatal(err)
	}
}	
