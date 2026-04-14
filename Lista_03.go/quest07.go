package main

import f "fmt"

func main() {

	var n, qtd, maior, menor, soma, soma_pares, qtd_pares, fin int
	var media, media_pares, percentagem float64
	primeiro := true

	for {
		f.Println("Digite um número inteiro positivo: ")
		f.Scan(&n)

		qtd++

		if primeiro {
			maior = n
			menor = n
			primeiro = false
		} else {
			if n > maior {
				maior = n
			}
			if n < menor {
				menor = n
			}
		}

		soma += n
		media = float64(soma) / float64(qtd)

		if n%2 == 0 {
			qtd_pares++
			soma_pares += n
			media_pares = float64(soma_pares) / float64(qtd_pares)
		}

		percentagem = (float64(qtd-qtd_pares) / float64(qtd)) * 100

		f.Println("Deseja finalizar a entrada?(Se sim, digite 30.000)")
		f.Scan(&fin)

		if fin == 30000 {
			break
		}

	}

	f.Println("\nA soma dos números digitados é: ", soma)
	f.Println("\nA quantidade de números digitados é: ", qtd)
	f.Printf("\nA média dos números digitados é: %.2f\n", media)
	f.Printf("\nO maior número é: %d\n", maior)
	f.Printf("\nO menor número é: %d\n", menor)
	f.Printf("\nA média dos números pares é: %.2f\n", media_pares)
	f.Printf("\nA percentagem dos números ímpares entre todos os números digitados é: %.2f%%\n", percentagem)

}
