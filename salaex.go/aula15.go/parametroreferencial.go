package main

import "fmt"

func main() {
	v := 0
	incrementa(&v)
	fmt.Println(v)
}
func incrementa(x *int) {
	*x++
}
