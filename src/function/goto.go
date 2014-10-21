package main
import (
	"fmt"
)
func myFunc() {
    i := 0
Here:   // label ends with ":"
    fmt.Println(i)
    i++
    goto Here   // jump to label "Here"
}
func main() {
	//error
	myFunc()
}