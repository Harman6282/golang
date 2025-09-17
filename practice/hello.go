package main

import "fmt"


func main() {

	language := struct {
		name string
		isGood bool
	} {"Golang", true}
  

	fmt.Println(language)

}
