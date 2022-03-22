package main

import (
	"fmt"
	// "strings"
)

// func main(){
// 	var evenNumbers = func(numbers ...int) []int{
// 		var result []int

// 		for _, v := range numbers {
// 			if v % 2== 0 {
// 				result = append(result, v)
// 			}
// 		}
// 		return result
// 	}

// 	var numbers = []int{4,93,77,10,52,22,34}

// 	fmt.Println(evenNumbers(numbers...))
// }

// func main(){

// 	var studentLists = []string{"Airell", "Nanda", "Mailo", "Schannel", "Marco"}

// 	var find = findStudent(studentLists)

// 	fmt.Println(find("airell"))
// }

// func findStudent(students []string) func(string) string {
// 	return func(s string)string {
// 		var student string
// 		var position int

// 		for i, v := range students {
// 			if strings.ToLower(v) == strings.ToLower(s){
// 				student = v
// 				position = i
// 				break
// 			}
// 		}
// 		if student == ""{
// 			return fmt.Sprintf("%s doesn't exists", s)
// 		}
// 		return fmt.Sprintf("We found %s at position %d", s, position+1)
// 	}
// }


func main() {
	var numbers  = []int{2,5,8,10,3,99,23}

	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 !=0
	})

	fmt.Println("Total Odd Numbers", find)
}

func findOddNumbers( numbers []int, callback func(int) bool) int{
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}