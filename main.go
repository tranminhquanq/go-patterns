package main

import (
	"github.com/tranminhquanq/go-patterns/singleton"
)

func main() {
	// // Singleton
	for i := 0; i < 10; i++ {
		go singleton.GetInstance()
	}
}
