package main

import (
	"fmt"
)
/*------------------[ Declaration variables ]-----------*/;
var i, j int = 1, 2
var c, python, java = true, false, "no!"

var i, j int = 1, 2
    c, python, java := true, false, "no!"

func main() {
	/*---------------Assign variable------------*/
	//Way 1
	/*var x string
	x = "xin chào"
	fmt.Println(x)*/
	//Way 2
	/*var x1 string = "Hello World"
	fmt.Println(x1)*/

	/*x1 := 5
	fmt.Println("x is", x1)*/

	// Multiple Variables
	/*var (
	   a = 5
	   b = 10
	   c = 15
	  )
	  fmt.Println("My dog's name is",a,b,c)*/

	//Constants
	/*const x string = "Hello World"
	fmt.Println(x)*/

	//String
	//Replace string
	/*s := "hello"
	  c := []byte(s)  // convert string to []byte type
	  c[0] = 'c'
	  s2 := string(c)  // convert back to string type
	  fmt.Printf("%s\n", s2)*/

	//Cut string
	/*s := "hello"
	  s = string(s[1:]) // you cannot change string values by index, but you can get values instead.
	  fmt.Printf("%s\n", s)*/

	//Mupltiline string
	/*m := `hello
	  world`
	  fmt.Printf("%s\n", m)
	*/

	/*-------- compare string ------------*/
	/*var x string = "hello"
	var y string = "hello"
	fmt.Println(x == y)*/

}
