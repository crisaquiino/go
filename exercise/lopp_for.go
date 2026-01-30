package main
import ( 
	"fmt"
)
func main () {
	x :=  7

	for {
		if x < 10 { 
			x++ 
			fmt.Println("teste")
	} else {
			fmt.Println("È maior")
			break
		}
	}
	fmt.Println("Loop break")
}	