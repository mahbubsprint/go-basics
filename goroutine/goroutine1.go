// package main

// import "fmt"

// func task(name string) {
//     fmt.Println("running:", name)
// }

// func main() {
//     go task("A")
//     go task("B")

//     // Without waiting, main may exit early.
//     fmt.Println("main done (maybe too soon)")
// }

package main

import (
    "fmt"
    "sync"
)

func work(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("worker", id, "start")
    // do work...
    fmt.Println("worker", id, "end")
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go work(i, &wg)
    }
    wg.Wait() // blocks until all Done() calls
    fmt.Println("all workers finished")
}