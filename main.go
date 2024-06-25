package main

import (
	"fmt"
	"github.com/extism/go-pdk"
)

//go:export greet
func Greet() int32 {
	name := pdk.InputString()
	pdk.OutputString("Hello, " + name)
	return 0
}

func main() {
    fmt.Println("Hello, World!!!!!!")
}
