a := 10
p := &a
a -= 4
*p = *p + 3
fmt.Println(a)
