package main

import f "fmt"

func main() {

	var n int

	for {
		f.Print("Digite um número inteiro (zero ou negativo para sair): ")
		f.Scan(&n)

		if n <= 0 {
			f.Println("Programa encerrado.")
			break
		}

		Quadrado_Perfeito := false

		for i := 1; i*i <= n; i++ {
			if i*i == n {
				Quadrado_Perfeito = true
				break
			}
		}

		if Quadrado_Perfeito {
			f.Printf("O número %d é um quadrado perfeito.\n", n)
		} else {
			f.Printf("O número %d NÃO é um quadrado perfeito.\n", n)
		}

	}
}
