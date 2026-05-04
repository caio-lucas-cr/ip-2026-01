package main

import f "fmt"

func main() {

	const totalPoltronas = 24
	var janela [totalPoltronas]int
	var corredor [totalPoltronas]int

	for {
		f.Println("\n--- SISTEMA DE VENDAS DE PASSAGENS ---")
		f.Println("1. Comprar na Janela")
		f.Println("2. Comprar no Corredor")
		f.Println("3. Finalizar")
		f.Print("Escolha uma opção: ")

		var opcao int
		f.Scan(&opcao)

		if opcao == 3 {
			f.Println("Sistema encerrado.")
			break
		}

		if opcao == 1 {
			venderPoltrona(&janela, "Janela")
		} else if opcao == 2 {
			venderPoltrona(&corredor, "Corredor")
		} else {
			f.Println("Opção inválida!")
		}

		if estaLotado(janela, corredor) {
			f.Println("\nAVISO: O ônibus está completamente lotado!")
			break
		}
	}
}

func venderPoltrona(setor *[24]int, nomeSetor string) {
	f.Printf("\n--- Poltronas Disponíveis na %s ---\n", nomeSetor)
	livres := 0
	for i := 0; i < 24; i++ {
		if setor[i] == 0 {
			f.Printf("[%d] ", i+1)
			livres++
		}
	}

	if livres == 0 {
		f.Printf("\nNão existem mais poltronas livres no setor %s!\n", nomeSetor)
		return
	}

	f.Print("\nQual poltrona deseja comprar? ")
	var escolha int
	f.Scan(&escolha)

	indice := escolha - 1
	if indice < 0 || indice >= 24 {
		f.Println("Número de poltrona inválido!")
	} else if setor[indice] == 1 {
		f.Println("Esta poltrona já está ocupada! Escolha uma da lista.")
	} else {
		setor[indice] = 1
		f.Println("Venda realizada com sucesso!")
	}
}

func estaLotado(janela [24]int, corredor [24]int) bool {
	for i := 0; i < 24; i++ {
		if janela[i] == 0 || corredor[i] == 0 {
			return false
		}
	}
	return true
}
