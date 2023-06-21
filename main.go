package main

import (
	"fmt"
	"lisp_by_go/lisp"
)

func main() {
	a := lisp.Int32{N: 12}
	name := "Go Developers"
	fmt.Println("Azure for", a, name)
}
