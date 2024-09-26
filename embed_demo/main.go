package main

import (
	_ "embed"
	"fmt"
)

//go:embed demo.txt
var data string

func main() {
	fmt.Println(string(data))
}
