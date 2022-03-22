package main

import (
	"fmt"
	// "math"
	"strings"
)

// func main(){
	// greet("Hacktiv8", 6)
	// greet("Hacktiv8", "jalan Sultan Iskandar Muda")

	// funtion with return
// 	var names = []string{"Hacktiv8", "Jakarta"}
// 	var printMsg = greet ("Heii", names)

// 	fmt.Println(printMsg)
//  }

// func greet(msg string, names []string) string{
	// fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
	// fmt.Println("Hello there! My Name is", name)
	// fmt.Println("I Live in", address)

// 	var joinStr = strings.Join(names,"")
// 	var result string = fmt.Sprintf("%s %s", msg, joinStr)

// 	return result
// }


// func main(){
// 	var diameter float64 = 15

// 	var area, circumference = calculate(diameter)

// 	fmt.Println("Area:", area)
// 	fmt.Println("Circumference:", circumference)	
// }


// func calculate(d float64) (float64, float64){

	//Menghitung Luas
	// var area float64 = math.Pi * math.Pow(d/2, 2)
	
	//Menghitung Keliling
// 	var circumference = math.Pi * d

// 	return area, circumference
// }

// func calculate( d float64) (area float64, circumference float64){
// 	//Menghitung luas
// 	area = math.Pi * math.Pow(d/2, 2)

// 	//Menghitung Keliling
// 	circumference = math.Pi * d

// 	return
// }

// func main(){
// 	studentLists := print("Hacktiv8", "Hacktivkidz", "Kode")

// 	fmt.Printf("%v", studentLists)
// }

// func print(names ...string) []map[string]string{
// 	var result []map[string]string

// 	for i, v := range names{
// 		key := fmt.Sprintf("student%d", i+1)
// 		temp := map[string]string{
// 			key: v,
// 		}
// 		result = append(result, temp)
// 	}
// 	return result
// }


func main() {
	numberLists := []int{1,2,3,4,5,6,7,8}

	result := sum(numberLists...)

	fmt.Println("Result:", result)
}

func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

func main(){
	profile("Airell", "Pasta", "Ayam geprek", "ikan roa", "sate")
}

func profile(name string, favFoods ...string){
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
