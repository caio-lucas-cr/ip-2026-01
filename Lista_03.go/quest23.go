package main

import f "fmt"

func main() {

	var N int
	var R float64

	f.Println("Digite a quantidade N (N > 0) de termos da série 1000/1 - 997/2 + 994/3 - 991/4 + ... :")
	f.Scan(&N)

	if N <= 0 {
		f.Println("Quantidade digitada inválida!")
		return
	}

	for i := 1; i <= N; i++ {
		if i%2 == 0 {
			R -= float64(1000-3*(i-1)) / float64(i)
		} else {
			R += float64(1000-3*(i-1)) / float64(i)
		}
	}

	f.Printf("O resultado da série com %d termos é: R = %.4f", N, R)

}
