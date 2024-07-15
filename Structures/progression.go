package main

import "fmt"

func main() {

	var (
		hold int;
		res int = 1;
	)
	fmt.Scan(&hold)

	for i := 1; i <= hold; i++ {
		res *= i;
	}
	fmt.Println(res)
}