package main

import "fmt"

func init() {
	fmt.Println("------------------------------------------")
	fmt.Println("AP4 - Torre de Hanói (Recursividade)")
	fmt.Println("------------------------------------------")
}

func main() {
	var n int = 2
	var origem, trabalho, destino string = "A", "B", "C"

	movimentosDiscos := hanoi(n, origem, trabalho, destino)
	fmt.Println("------------------------------------------")
	fmt.Printf("O número de movimentos para %d discos é %d\n", n, movimentosDiscos)
	fmt.Println("------------------------------------------")
}

func hanoi(n int, pinoOrigem, pinoTrabalho, pinoDestino string) int {
	var contador int = 0
	if n == 0 {
		return 0
	}
	contador += hanoi(n-1, pinoOrigem, pinoDestino, pinoTrabalho)
	fmt.Printf("Disco %d da torre %s para torre %s\n", n, pinoOrigem, pinoDestino)
	contador++
	contador += hanoi(n-1, pinoTrabalho, pinoOrigem, pinoDestino)
	return contador
}
