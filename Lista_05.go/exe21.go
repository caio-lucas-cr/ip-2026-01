package main

import f "fmt"

func main() {
	var codigo int
	var vetor [10]float64

	f.Print("Digite o código (0:Sair, 1:Direta, 2:Inversa): ")
	f.Scan(&codigo)

	if codigo == 0 {
		f.Println("Programa encerrado.")
		return
	}

	if codigo != 1 && codigo != 2 {
		f.Println("Erro: Código inválido! Use apenas 0, 1 ou 2.")
		return
	}

	f.Println("Digite os 10 números reais:")
	for i := 0; i < 10; i++ {
		f.Printf("Vetor[%d]: ", i)
		f.Scan(&vetor[i])
	}

	if codigo == 1 {
		f.Println("\n--- Ordem Direta ---")
		for i := 0; i < 10; i++ {
			f.Printf("%.2f ", vetor[i])
		}
	} else if codigo == 2 {
		f.Println("\n--- Ordem Inversa ---")
		for i := 9; i >= 0; i-- {
			f.Printf("%.2f ", vetor[i])
		}
	}
	f.Println()
}