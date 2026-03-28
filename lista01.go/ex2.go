package main

import f "fmt"

func main() {

	//Algoritmo "Arrecadação de jogos"

	var ingressos float32
	var perc_popular, perc_geral, perc_arquibancadas, perc_cadeiras, renda float32

	//Entrada
	f.Println("Digite a quantidade de ingressos: ")
	f.Scan(&ingressos)

	f.Println("Digite o percentagem de pessoas na categoria popular: ")
	f.Scan(&perc_popular)

	f.Println("Digite o percentagem de pessoas na categoria geral: ")
	f.Scan(&perc_geral)

	f.Println("Digite o percentagem de pessoas na categoria arquibancadas: ")
	f.Scan(&perc_arquibancadas)

	f.Println("Digite o percentagem de pessoas na categoria cadeiras: ")
	f.Scan(&perc_cadeiras)

	//Processamento
	renda = (perc_popular*1 + perc_geral*5 + perc_arquibancadas*10 + perc_cadeiras*20) * ingressos / 100

	//Saída
	f.Println("A renda do jogo é: ", renda)

}
