package main

import (
	"fmt"
	// "time"
)

// func main() {
// 	fmt.Println(1 )
// 	fmt.Println(true)
// 	fmt.Print(12.33/ 4.3)
// }

// ? variables

func main() {

	// * Day 1
	// var name string = "golang"
	// var name = "Harman"

	// age := 23
	// var name string
	// name = "hello"

	// fmt.Println(name)

	//* Day 2
	// const age = 19

	// const (
	// 	port = 5000
	// 	host = "localhost"
	// )

	// fmt.Println(port, host)

	// ? Loop
	// while loop
	//   i := 1
	//   for i <= 5{
	// 	fmt.Println(i)
	// 	i = i +1
	//   }

	// ? for loop
	// for i:= 1; i <= 5 ; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i:= range 4{
	// 	fmt.Println(i)
	// }

	// ?  If else (no ternary operator)

	// age := 14

	// if age > 18 {
	// 	fmt.Println("person is adult")
	// } else if age >= 12 {
	//    fmt.Println("person is teen")
	// }else {
	// 	fmt.Println("person is not adult")
	// }

	// var role = "admin"
	// var authorised = true

	// if role == "admin" && authorised{
	// 	fmt.Println("allowed")
	//   }

	// if age := 19; age >= 18 {
	// 	fmt.Println("person is adult")
	// }

	//? switch statement

	// i := 2

	// switch i{
	// case 1:
	//   fmt.Println("one")

	// case 2:
	//   fmt.Println("one")

	// case 3:
	//   fmt.Println("one")

	// default:
	// 	fmt.Println("default")
	// }

	// ? multiple condition switch

	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("weekend")
	// default:
	// 	fmt.Println("Its work day")
	// }

	//?  type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("integer")

		case string:
			fmt.Println("string")

		case bool:
			fmt.Println("boolean")

		default:
			fmt.Println("other", t)

		}

	}

	whoAmI("")

}
