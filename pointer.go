package main

import (
	"fmt"
	// "strings"
)

// func main(){
// 	var firstNumber *int
// 	var secondNumber *int
// 	_, _ = firstNumber, secondNumber
// }

// func main(){
// 	var firstNumber int = 4
// 	var secondNumber *int = &firstNumber

// 	fmt.Println("firstNumber (value) :", firstNumber)
// 	fmt.Println("firstNumber (memori address :", &firstNumber)

// 	fmt.Println("secondNumber (value) :", *secondNumber)
// 	fmt.Println("secondNumber (memori address :", secondNumber)
// }

// func main(){
// 	var firstPerson string ="John"
// 	var secondPerson *string = &firstPerson

// 	fmt.Println("firstPerson (value) :", firstPerson)
// 	fmt.Println("firstPerson (memori address) :", &firstPerson)
// 	fmt.Println("secondNumber (value) :", *secondPerson)
// 	fmt.Println("secondNumber (memori address :", secondPerson)

// 	fmt.Println("\n", strings.Repeat("#", 30), "\n")

// 	*secondPerson = "Doe"

// 	fmt.Println("firstPerson (value) :", firstPerson)
// 	fmt.Println("firstPerson (memory address) :", &firstPerson)
// 	fmt.Println("secondPerson (value) :", *secondPerson)
// 	fmt.Println("secondPerson (memory address) :", secondPerson)

// }


func main(){
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)
}

func changeValue(number *int){
	*number = 20
}
