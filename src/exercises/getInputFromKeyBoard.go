package main

import (
	"fmt"
)

func main() {
	var response int
	fmt.Println("This program will activate SkyNet worldwide, are you sure about this?\r\n")

	//fmt.Scanf("%c", &response) //<--- here
	/*	response, err := fmt.Scanf("%s\n", &response)
		if err != nil {
			fmt.Println("error", err)
		}*/
	switch response {
	default:
		fmt.Println("SkyNet launch aborted!")
	case 'y':
		fmt.Println("SkyNet launched")
	case 'Y':
		fmt.Println("SkyNet launched")
	}
}
