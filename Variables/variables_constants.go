package main

import "fmt"

const Pi = 3.14

func main() {
	const world = "disney"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}