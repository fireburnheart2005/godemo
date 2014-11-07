package main

import "fmt"

func main() {
	/*------------------[ Way 1 ]-----------*/
	/*sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)*/

	/*------------------[ Way 2 ]-----------*/
	/*i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}*/

	/*------------------[ demo ]-----------*/
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

}
