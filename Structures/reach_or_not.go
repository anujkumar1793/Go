package main

import "fmt"

func main() {
    var money int
    fmt.Scan(&money)
    if money > 1000 {
        fmt.Println("Apple")
    } else if money >= 500 && money <= 1000 {
        fmt.Println("Samsung")
    } else {
        fmt.Println("Nokia с фонариком")
    }
}