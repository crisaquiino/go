package main

import (
	"fmt"
)

func main (){

	sabores := []string{"abacaxi", "uva", "goiaba", "alface", "tomate"}
	fatia := sabores[0:]
	fmt.Println(fatia) 
}

