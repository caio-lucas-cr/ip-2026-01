package main

import (
	"fmt"
	"sort"
)

type Empregado struct {
	ID    int
	Meses int
}

func main() {
	var empregados []Empregado

	fmt.Println("Insira o número do empregado e os meses de trabalho (0 0 para encerrar):")

	for i := 0; i < 100; i++ {
		var id, meses int
		fmt.Printf("Entrada %d: ", i+1)
		fmt.Scan(&id, &meses)

		if id == 0 && meses == 0 {
			break
		}

		empregados = append(empregados, Empregado{ID: id, Meses: meses})
	}

	if len(empregados) == 0 {
		fmt.Println("Nenhum dado foi inserido.")
		return
	}

	// Ordenação: usamos o pacote 'sort' para ordenar do menor tempo para o maior
	// Menos meses = mais recente
	sort.Slice(empregados, func(i, j int) bool {
		return empregados[i].Meses < empregados[j].Meses
	})

	fmt.Println("\n--- Os 3 empregados mais recentes ---")
	
	// Definimos o limite para 3 ou o tamanho total (caso haja menos de 3)
	limite := 3
	if len(empregados) < 3 {
		limite = len(empregados)
	}

	for i := 0; i < limite; i++ {
		fmt.Printf("%dº mais recente: ID %d (%d meses de empresa)\n", 
			i+1, empregados[i].ID, empregados[i].Meses)
	}
}