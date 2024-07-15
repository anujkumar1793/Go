package main
import "fmt"

func maximum(a int, b int) int{
	if a > b {
		return a
	} else {
		return b
	}
}

func calc(a int) (int, int) {
	return a * 2, a * a
}

func isEven(a int) bool{

	if a % 2 == 0{
		return true
	} else {
		return false
	}
}


func double(a int) {
    fmt.Println(a*2) 
} 


func squareSumm(a int, b int) int {
	var ss int;

	for i := a; i <= b; i++ {
		ss += i * i
	}
	return ss
}


func main() {
    for i := 4; i > 0; i-- {
        defer double(i)
    }
}


func oneOrTwo(a int, b int, s string) (int, string){

	if s == "two"{
		return b, s
	} else if s == "one"{
		return a, s
	} else {
		return 0, s
	}
}



func marsAge(age int) int {
	var mage int;

	mage = age * 365 / 687 

	return mage
}
