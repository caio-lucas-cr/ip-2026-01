package main

import f "fmt"

func main() {

	fibonacci := make([]uint64, 50)

	fibonacci[0] = 1
	fibonacci[1] = 1

	for i := 2; i < len(fibonacci); i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	f.Printf("Os primeiros 50 termos da série de Fibonacci:\n")
	f.Printf("\nSequência de Fibonacci: %d", fibonacci[0])

	for i:=1; i<len(fibonacci); i++{
		f.Printf(",%d", fibonacci[i])
	}
	
}