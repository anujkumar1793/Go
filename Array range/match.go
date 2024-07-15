package main

import "fmt"

func main() {
    var res string
    var total int
    var draw int

    fmt.Scanln(&res)
    
    results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
    results = append(results, res)

    for _, v := range results {
        if v == "w" {
            total += 1
        }
        if v == "d"{
            draw += 1    
        }
    } 

    fmt.Println(total * 3 + draw )
}
