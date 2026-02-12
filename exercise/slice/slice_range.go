package main

import (
	"fmt"
)

func main(){
	slice := []string{"banana", "maça", "jaca", "uva"}
	for indice, valor := range slice {
		fmt.Println("No indice", indice, "temos o valor", valor)
	}
	slice[3] ="melancia"

	for indice, valor := range slice {
		fmt.Println("No indice", indice, "temos valor", valor)
	}
}


// slice voce consegue adicionar vvalores as variaveis o que nao e permitido no array