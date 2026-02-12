package main

import (
	"fmt"
)

func main (){
	slice := make([]int, 5, 10)
	slice[0], slice[1], slice[2], slice[3] = 1, 2, 3, 4
	//slice[5] = append(slice, 10)

	fmt.Println(slice, len(slice), len(slice), cap(slice))

	
}