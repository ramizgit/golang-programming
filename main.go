package main

import (
	"fmt"

	"example.com/hello/hello"

	"example.com/hello/array"
)

func main() {
	greeting := "This is the starting point of golang execution."

	fmt.Print(greeting)

	//call hello package
	fmt.Print(hello.GetMessage())

	fmt.Print(hello.Hello())

	fmt.Print(hello.Quote())

	//call array package
	array.GetFruits()

}
