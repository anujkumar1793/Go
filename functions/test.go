package main

import "fmt"

var x = 13

func one() {
    x += 1
}

func two(x int) {
    x += 3
}

func main() {
    one()
    two(x)
    fmt.Println(x)
}