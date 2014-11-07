package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float32
	fmt.Scanf("%f\n", &input)
	n, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(n, err)
	}
	output := input * 2
	fmt.Printf("%.2f", output)
}
