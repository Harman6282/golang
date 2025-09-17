package main

import (
	"fmt"
)

// func processNum(numChain chan int) {

// 	for num := range numChain {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second * 1)
// 	}

// }

// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2

// 	result <- numResult

// }

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Processing...")

}

func main() {

	done := make(chan bool)
	go task(done)

	<- done

	// result := make(chan int)

	// go sum(result, 3, 6)

	// res := <- result

	// fmt.Println(res)

	// numChain := make(chan int)

	// go processNum(numChain)

	// for {
	// 	numChain <- rand.Intn(100)
	// }

	// messageChan := make(chan string)

	// messageChan <- "ping"

	// msg := <- messageChan

	// fmt.Println(msg)
}
