package main

import (
	f "fmt"
)

func main() {

	const totalContas = 10
	var codigos [totalContas]int
	var saldos [totalContas]float64

	f.Println("--- Cadastro Inicial de 10 Contas ---")
	
	for i := 0; i < totalContas; i++ {
		for {
			var cod int
			f.Printf("Digite o código da conta %d: ", i+1)
			f.Scan(&cod)

			repetido := false
			for j := 0; j < i; j++ {
				if codigos[j] == cod {
					repetido = true
					break
				}
			}

			if repetido {
				f.Println("Erro: Já existe uma conta com este código. Tente novamente.")
			} else {
				codigos[i] = cod
				f.Printf("Digite o saldo inicial da conta %d: ", cod)
				f.Scan(&saldos[i])
				break 
			}
		}
	}

	for {
		f.Println("\n--- CONTROLE BANCÁRIO ---")
		f.Println("1. Efetuar depósito")
		f.Println("2. Efetuar saque")
		f.Println("3. Consultar ativo bancário")
		f.Println("4. Finalizar programa")
		f.Print("Escolha uma opção: ")

		var opcao int
		f.Scan(&opcao)

		if opcao == 4 {
			f.Println("Encerrando o sistema...")
			break 
		}

		switch opcao {
		case 1: 
			var buscaCod int
			var valor float64
			f.Print("Digite o código da conta para depósito: ")
			f.Scan(&buscaCod)

			indice := -1
			for i := 0; i < totalContas; i++ {
				if codigos[i] == buscaCod {
					indice = i
					break
				}
			}

			if indice == -1 {
				f.Println("Conta não encontrada.")
			} else {
				f.Print("Valor a depositar: ")
				f.Scan(&valor)
				saldos[indice] += valor
				f.Println("Depósito realizado com sucesso!")
			}

		case 2: 
			var buscaCod int
			var valor float64
			f.Print("Digite o código da conta para saque: ")
			f.Scan(&buscaCod)

			indice := -1
			for i := 0; i < totalContas; i++ {
				if codigos[i] == buscaCod {
					indice = i
					break
				}
			}

			if indice == -1 {
				f.Println("Conta não encontrada.")
			} else {
				f.Print("Valor a sacar: ")
				f.Scan(&valor)
				if saldos[indice] >= valor {
					saldos[indice] -= valor
					f.Println("Saque efetuado com sucesso!")
				} else {
					f.Println("Saldo insuficiente.")
				}
			}

		case 3:
			somaAtivo := 0.0
			for i := 0; i < totalContas; i++ {
				somaAtivo += saldos[i]
			}
			f.Printf("Ativo total do banco: R$ %.2f\n", somaAtivo)

		default:
			f.Println("Opção inválida! Tente novamente.")
		}
	}
}