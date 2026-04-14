package main

import f "fmt"

func main() {

	var N, al_apro, al_ex, al_rep, qtd int
	var n1, n2, med_al, soma_med_al, med_class float64
	var mensagem string

	f.Println("Digite a quantidade de alunos: ")
	f.Scan(&N)

	for i := 1; i <= N; i++ {
		f.Printf("\n Aluno %d \n", i)
		f.Println("Digite a 1ª nota: ")
		f.Scan(&n1)

		f.Println("Digite a 2ª nota: ")
		f.Scan(&n2)

		med_al = (n1 + n2) / 2

		f.Printf("\nA média do aluno %d é: %.2f\n", i, med_al)
		qtd++

		if med_al >= 0 && med_al <= 3 {
			al_rep++
			mensagem = "Reprovado"
			f.Println("O aluno está", mensagem)
		} else if med_al > 3 && med_al < 7 {
			al_ex++
			mensagem = "Exame"
			f.Println("O aluno ficou de", mensagem)
		} else if med_al >= 7 {
			al_apro++
			mensagem = "Aprovado"
			f.Println("O aluno está", mensagem)
		}

		soma_med_al += med_al

		med_class = soma_med_al / float64(qtd)

	}

	f.Printf("\nA quantidade de alunos aprovados é: %d \n", al_apro)
	f.Printf("\nA quantidade de alunos de exame é: %d \n", al_ex)
	f.Printf("\nA quantidade de alunos reporvados é: %d \n", al_rep)
	f.Printf("\nA média da classe é: %.2f", med_class)

}
