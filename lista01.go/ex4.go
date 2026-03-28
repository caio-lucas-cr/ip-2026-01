package main

import f "fmt"

func main() {

	//Algoritmo "Consumo de energia"

	var salario, kW, val_kW, val_cons, val_desc float32

	//Entrada
	f.Println("Digite o valor do salário mínimo: ")
	f.Scan(&salario)

	f.Println("Digite a quantidade de kW consumido por residência: ")
	f.Scan(&kW)

	//Processamento
	val_kW = salario * 0.007
	val_cons = val_kW * kW
	val_desc = val_cons * 0.9

	//Saída
	f.Println("Custo por kW: R$ ", val_kW)
	f.Println("Custo do consumo: R$ ", val_cons)
	f.Println("Custo com desconto: R$ ", val_desc)

}
