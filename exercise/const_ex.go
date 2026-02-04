package main
import (
	"fmt"
)

func main (){
	x := 100
	fmt.Printf("%d, %#x, %b\n", x, x, x)
	y := x >> 1
	fmt.Printf("%d, %#x, %b", y, y, y)
}

// atribuir valores, decimal, hexa e binario