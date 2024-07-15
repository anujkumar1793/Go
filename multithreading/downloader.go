package main
import "fmt"

func download(v int, ch chan int){
	csum := 0
	for i := 0; i <= v; i++ {
		csum += i
	}

	ch <- csum
} 
 

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)
 
    var s1, s2, s3 int
    fmt.Scanln(&s1)
    fmt.Scanln(&s2) 
    fmt.Scanln(&s3)

    go download(s1, ch1)
    go download(s2, ch2)
    go download(s3, ch3)

    value1 := <- ch1
	value2 := <- ch2
	value3 := <- ch3
	fmt.Println(value1 + value2 + value3)
	
}