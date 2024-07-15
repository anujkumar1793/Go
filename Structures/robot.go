package main

import "fmt"

func main() {
	var a, b, c string

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	
	dict := make(map[string]string)

	dict["0"] = "Ноль"
	dict["1"] = "Один"
	dict["2"] = "Два"
	dict["3"] = "Три"
	dict["4"] = "Четыре"
	dict["5"] = "Пять"
	dict["6"] = "Шесть"
	dict["7"] = "Семь"
	dict["8"] = "Восемь"
	dict["9"] = "Девять"
	dict["10"] = "Десять"

	fmt.Println(dict[a])
	fmt.Println(dict[b])
	fmt.Println(dict[c])
}
