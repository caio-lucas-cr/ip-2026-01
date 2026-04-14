package main

import f "fmt"

func main() {

	var N int

	f.Println("Insira a quantidade de termos da série: ")
	f.Scan(&N)

	x := 1

	f.Println("\nA série é:")
	f.Printf("%d", x)

	for i := 2; i <= N; i++ {
		k := i * i
		f.Printf(", %d", k)
	}
}
