package main

import f "fmt"

func main() {

	f.Println("Exibindo os índices acima da diagonal principal da matriz 10x10 [linha, coluna]:")

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i != j && i < j {
				f.Printf("[%d,%d]\t", i, j)
			}
		}
		f.Println()
	}

}
