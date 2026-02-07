package main

import (
	"fmt"
)

func main(){
	array := [5]int{1,2,3,4,5}
	fmt.Println(array)
	slice := []int{1,2,3,4,5}
	fmt.Println(slice)


		slice2 := append(slice, 6)

	fmt.Println(slice2)
	fmt.Println(slice2[2])
}


// slice voce consegue adicionar vvalores as variaveis o que nao e permitido no array