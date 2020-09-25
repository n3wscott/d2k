package main

import (
	"os"
	"tableflip.dev/d2k/pkg/demo"
)

func main() {
	demo.Hello(os.Getenv("HELLO"))
}
