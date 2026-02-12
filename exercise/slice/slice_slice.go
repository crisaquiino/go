package main

import (
	"fmt"
)

func main (){

	sabores := []string{"abacaxi", "uva", "goiaba", "alface", "tomate"}
//	fatia := sabores[:]
//	fmt.Println(fatia) 

	sabores = append(sabores[:1], sabores[2:]...)
	fmt.Println(sabores)

}

