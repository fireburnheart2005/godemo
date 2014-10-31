package main

import (
	"fmt"
)

/*var x string = "Hello World"

func main() {
	fmt.Println(x)
}
func f() {
	fmt.Println(x)
}*/

//Error scope
func main() {
	var x string = "Hello World"
	fmt.Println(x)
}

func f() {
	fmt.Println(x)
}
