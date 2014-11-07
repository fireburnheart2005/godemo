package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func split2(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var i, j int = 1, 2
var c, python, java = true, false, "no!"

func main() {
	//fmt.Println(split(17))
	fmt.Println(i, j, c, python, java)
}
