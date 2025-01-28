package main 

import (
	"fmt"
)

// Go variable type 
// int, float32, string, boll and others

func main(){
	//Syntax
	syntax := "var variablename type = value"
	fmt.Println(syntax)

	//Declaration with initial value
	var student1 string = "Jonh" //type is string
	var student2 = "Jane"
	X := 2

	fmt.Println("\n")
	fmt.Println(student1)
    fmt.Println(student2)
    fmt.Println(X)

	fmt.Println("\n")
	
	var a string
	var b int 
	var c bool

	fmt.Printf("%q\n", a)
	fmt.Println(b)
	fmt.Println(c)

	a = "O Neymar tá de volta!"

	fmt.Println(a)
}