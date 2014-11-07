package main

import (
	"fmt"
)

func func_switch1() {
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}
}

func func_switch2() {
	integer := 6
	switch integer {
	case 4:
		fmt.Println("integer <= 4")
		fallthrough
	case 5:
		fmt.Println("integer <= 5")
		fallthrough
	case 6:
		fmt.Println("integer <= 6")
		fallthrough
	case 7:
		fmt.Println("integer <= 7")
		fallthrough
	case 8:
		fmt.Println("integer <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
func main() {
	//Normat switch break
	func_switch1()
	//Continue using thefallthrough not break
	//func_switch2()
}
