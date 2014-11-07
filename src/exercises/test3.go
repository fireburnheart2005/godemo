package main

import "fmt"

func main() {
	var i, j int
	_, err := fmt.Scanf("%d %d\n", &i, &j)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(cf(i, j))
}
func cf(n int, c int) (r int) {
	r = 1
	for i := 0; i < c; i++ {
		r = r * n
	}
	return
}
