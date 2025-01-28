package main 

import (
	"fmt"
)

func main(){
	ex1()
	ex2()
	ex3()
	ex4()
}

func ex1(){
	var myNum int = 50
	fmt.Println(myNum)
}

func ex2(){
	var myWord = "hello!"
	fmt.Println(myWord)
}

func ex3(){
	var x = 5
	var y = 10
	var z = x + y
	fmt.Println(z)
}

func ex4(){
	var x, y, z = 5, 10, 15
	fmt.Println(x, y, z)

	fmt.Printf("Value of X: %v\n", x)
	fmt.Printf("Value of X: %v\n", y)
	fmt.Printf("Value of X: %v\n", z)
}
