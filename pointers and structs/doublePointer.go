func doublePointer(x *int) int {
	*x = *x * 2 
	return *x
}