package main

import "fmt"

func main() {
	//Invoke the sayHelloAndReturnNothingness function without any parameters
	sayHelloAndReturnNothingness()
	a := myNumber()
	fmt.Println(a)
}
func sayHelloAndReturnNothingness() {
	fmt.Println("I'm a function in GO!")
}

func myNumber() int {
	return 42
}
