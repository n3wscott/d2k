package main

import (
	"fmt"
	"os"
	"tableflip.dev/d2k/pkg/demo"
)

func main() {
	fmt.Println(demo.Hello(os.Getenv("HELLO")))
}
