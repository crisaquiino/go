package main
import (
	"fmt"
)

func main (){
	x := 5

	switch {
	case x == 5:
		fmt.Println("x = 5 primeiro")
	case x == 5:
		fmt.Println("x e igual a cinco")
	case x == 5:
		fmt.Println("menor que 5")
	default:
		fmt.Println("Valor 5")
	}
}