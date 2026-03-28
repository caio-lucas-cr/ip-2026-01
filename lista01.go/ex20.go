package main

import f "fmt"

func main() {

	//Algoritmo "Tempo em segundos"

	var h, min, s, seg int

	//Entrada
	f.Println("Digite as horas, minutos e segundos: ")
	f.Scan(&h, &min, &s)

	//Processamento
	seg = h*60*60 + min*60 + s

	//Saída
	f.Println("A quantidade de segundos é: ", seg)

}
