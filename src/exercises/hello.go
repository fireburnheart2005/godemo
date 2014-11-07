package main

import "fmt"

func main() {
	fmt.Print("Please input your name : ")
	var input string
	/*------------------[ way 1 ]-----------*/
	fmt.Scanf("%s", &input)
	/*------------------[ way 2 ]-----------*/
	fmt.Scanln(&input)
	//output := input * 2
	fmt.Printf("%s", "Hello"+input)
}
