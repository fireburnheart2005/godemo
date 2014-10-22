package main
import "fmt"

// return results of A + B and A * B
func SumAndProduct(A, B int) (int, int, float32) {
return A+B, A*B,float32(A)/float32(B)
}

func SumAndProduct2(A, B int) (a int, b int, c float32) {
	a = A + B;
	b = A*B;
	c = float32(A)/float32(B);
    //return A+B, A*B,float32(A)/float32(B)
    return
}

func main() {
    x := 3
    y := 4

    xPLUSy, xTIMESy,xDIV := SumAndProduct(x, y)

    fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
    fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
    fmt.Printf("%d / %d = %.2f\n", x, y, xDIV)
}