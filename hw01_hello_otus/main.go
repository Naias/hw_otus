package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	got := reverse.String("Hello, OTUS!")
	fmt.Printf("%v\n", got)
}