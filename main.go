package main

import (
	"fmt"

	"github.com/nari-z/go-generics-example/sample"
)

func main() {
	s := sample.New[string]()
	s.Insert("hello")
	s.Insert("world")

	fmt.Print("s: %v", s)

	n := sample.New[int]()
	n.Insert(3)
	n.Insert(8)
	n.Insert(8)

	fmt.Print("n: %v", n)
}
