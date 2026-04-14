package main

import f "fmt"

func main() {

	var S float64

	f.Println("A soma dos 20 primeiros termos de da série 100/0! + 99/1! + 98/2! 97/3! + ... é: ")

	S = 100
	k := 1

	for i := 1; i <= 19; i++ {
		k = k * i
		S += float64(100-i) / float64(k)
	}

	f.Printf("S = %.4f", S)

}
