package main

import (
	f "fmt"
)

func main() {
	const totalAlunos = 15
	var frequenciaAbsoluta [11]int
	var nota int

	f.Printf("Insira as %d notas (de 0 a 10):\n", totalAlunos)

	for i := 0; i < totalAlunos; i++ {
		f.Printf("Nota do aluno %d: ", i+1)
		f.Scan(&nota)

		if nota >= 0 && nota <= 10 {
			frequenciaAbsoluta[nota]++
		} else {
			f.Println("Nota inválida! Por favor, insira um valor entre 0 e 10.")
			i-- 
		}
	}

	f.Println("\n--- TABELA DE FREQUÊNCIAS ---")
	f.Printf("%-6s | %-10s | %-10s\n", "Nota", "F. Absoluta", "F. Relativa")
	f.Println("-------------------------------------------")

	for notaItem := 0; notaItem <= 10; notaItem++ {
		abs := frequenciaAbsoluta[notaItem]
		relativa := float64(abs) / float64(totalAlunos)

		f.Printf("  %2d   |     %2d      |     %.2f (%.1f%%)\n", notaItem, abs, relativa, relativa*100)
	}
}