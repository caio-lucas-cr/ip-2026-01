package main

import f "fmt"

func main() {

	var S float64

	f.Println("A soma S=37*38/1 + 36*37/2 + 35*36/3 + ... + 1*2/37 é: ")

	for i := 1; i <= 37; i++ {
		S += float64((38-i)*(39-i)) / float64(i)
	}

	f.Printf("S = %.2f", S)

}
