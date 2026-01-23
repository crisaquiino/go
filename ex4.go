package main

import (
	"fmt"
)
type hotdog int
var x hotdog

func main () {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)	
}


// - Crie um tipo subjacente deve ser int
// - Crie uma variavel para este tipo, identificador "x", uitilizando a palavra chave var
// - Na funcao main
