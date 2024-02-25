package main

import (
	"fmt"
)

type Page struct {
	Title string
	Body []byte
}

func main() {
	fmt.Println("Hello, Wiki!")
}