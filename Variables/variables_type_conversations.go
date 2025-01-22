package main 

import (
	"fmt"
	"math"
)

func main(){
	var x, y int = 3, 4

	fmt.Printf("Tipo de x = %T, Valor de x = %v\n",x ,x)
	fmt.Printf("Tipo de y = %T, Valor de y = %v\n",y ,y)

	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	
	fmt.Printf("Tipo de f = %T, Valor de f = %v\n",f ,f)
	fmt.Printf("Tipo de z = %T, Valor de z = %v\n",z ,z)

}