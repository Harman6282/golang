package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // mark as done when finished
    fmt.Printf("Worker %d starting\n", id)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)         // add one goroutine
        go worker(i, &wg) // start goroutine
    }

    wg.Wait() // wait for all goroutines
    fmt.Println("All workers finished")
}
