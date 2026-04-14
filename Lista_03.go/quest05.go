package main

import f "fmt"

func main() {

	var alt, peso, porc, soma_alt, m float64
	var id, cont, id_10_e_20, peso_me40, id_ma50, qtd_total int

	for {
		f.Println("Insira a idade, altura e peso da pessoa, respectivamente: ")
		f.Scan(&id, &alt, &peso)

		qtd_total++

		if id > 50 {
			id_ma50++
		}

		if id > 10 && id < 20 {
			soma_alt += alt
			id_10_e_20++
		}

		if peso < 40 {
			peso_me40++
		}

		f.Println("Deseja continuar inserindo dados? (Digite 1 para sim, qualquer outro valor encerrará o processo)")
		f.Scan(&cont)

		if cont != 1 {
			break
		}

	}

	if id_10_e_20 > 0 {
		m = soma_alt / float64(id_10_e_20)
	}

	if qtd_total > 0 {
		porc = float64(peso_me40) * 100 / float64(qtd_total)

	}

	f.Printf("\nA quantidade de pessoas com idade superior a 50 anos é: %d\n", id_ma50)
	f.Printf("A média das alturas das pessoas com idade entre 10 e 20 anos é: %.2f\n", m)
	f.Printf("A porcentagem de pessoas com peso inferior a 40 quilos é: %.2f%%\n", porc)

}
