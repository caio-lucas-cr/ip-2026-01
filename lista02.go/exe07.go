package main

import f "fmt"

func main() {

	var A, B, C, MENOR, INTER, MAIOR int

	f.Println("Digite números distintos, denomidados A, B e C, respectivamente: ")
	f.Scan(&A, &B, &C)

	if A < B && A < C && B < C {
		MENOR = A
		INTER = B
		MAIOR = C
		f.Println(MENOR, ",", INTER, ",", MAIOR)
	} else if A < C && A < B && C < B {
		MENOR = A
		INTER = C
		MAIOR = B
		f.Println(MENOR, ",", INTER, ",", MAIOR)
	} else if B < A && B < C && A < C {
		MENOR = B
		INTER = A
		MAIOR = C
		f.Println(MENOR, ",", INTER, ",", MAIOR)
	} else if B < C && B < A && C < A {
		MENOR = B
		INTER = C
		MAIOR = A
		f.Println(MENOR, ",", INTER, ",", MAIOR)
	} else if C < B && C < A && B < A {
		MENOR = C
		INTER = B
		MAIOR = A
		f.Println(MENOR, ",", INTER, ",", MAIOR)
	} else if C < A && C < B && A < B {
		MENOR = C
		INTER = A
		MAIOR = B
		f.Println(MENOR, ",", INTER, ",", MAIOR)
	}

}
