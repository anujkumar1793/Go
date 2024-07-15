package main

import "fmt"

func some(b *int) {
    *b = *b - 3
}

func main() {
    a := 11
    some(&a)
    fmt.Println(a)
}