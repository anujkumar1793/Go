package main

import "fmt"

func main() {
    var height int
    fmt.Scan(&height)
    if height > 185 {
        fmt.Println("Высокий рост")
    } else if height >= 170 && height <= 185 {
        fmt.Println("Средний рост")
    } else {
        fmt.Println("Небольшой рост")
    }
}
