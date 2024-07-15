package main

import "fmt"

func test(s string) string {
    return s
}


func main(){

	var str1 string;
	var str2 string;
	fmt.Scan(&str1, &str2)
	fmt.Println(test(str1), test(str2))
}

// func myStr(s1 string, s2 string) (string, string) {

// 	return s1, s2
// }




