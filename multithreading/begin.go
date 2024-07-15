//Многопоточность - это создание нескольких процессов, выполняющихся независимо друг от друга.

package main

import (
    "fmt"
    "time"
    )

func out(from, to int, ch chan bool) {
    for i:=from; i<=to; i++ {
        time.Sleep(50 * time.Millisecond)
        fmt.Println(i)
    }
    ch <- true
}

func main() {
    ch := make(chan bool)

    go out(0, 5, ch)
    go out(6, 10, ch)

    <-ch
    <-ch
}
