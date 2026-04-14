package main

import f "fmt"

func main() {

	var X, S float64

	f.Println("Digite um número X: ")
	f.Scan(&X)

	S = X
	k := 1

	for i := 1; i <= 19; i++ {
		k = k * i

		if i%2 == 0 {
			S += X / float64(k)
		} else {
			S -= X / float64(k)
		}
	}

	f.Printf("\nO somatório S é: %.4f\n", S)

}
