package main

import f "fmt"

func main() {

	//Algoritmo: "Aprovado ou Reprovado"

	var nota1, nota2, nota3, media float32

	//Entrada
	f.Println("Digite sua nota 1: ")
	f.Scan(&nota1)

	f.Println("Digite sua nota 2: ")
	f.Scan(&nota2)

	f.Println("Digite sua nota 3: ")
	f.Scan(&nota3)

	//Processamento
	media = (nota1 + nota2 + nota3) / 3

	f.Println("A média do aluno é: ", media)

	//Saída
	if media >= 6 {
		f.Println("Status: Aprovado")
	} else {
		f.Println("Status: Reprovado")
	}

}
