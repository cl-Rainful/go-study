package main

import (
	"fmt"
	"os"

	"rainful.link/hello"
)

func main() {
	// fmt.Print("hello world")
	// print("hello world")

	// hello.Print()

	hello.Print()

	// f, err := os.Open("./hello/hello.go")
	// f.Close()
	// fmt.Printf("err: %v\n", err)

	f, _ := os.Open("./hello/hello.go")
	fmt.Println(f.Name())
}
