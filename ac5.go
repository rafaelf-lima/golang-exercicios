/* Dado um array de números inteiros positivos, considerado ordenado crescentemente, e um valor
alvo, encontre um par de números no array cuja soma seja igual ao valor alvo. Se nenhum par for
encontrado, retorne um valor (-1, -1) indicando que nenhum par foi encontrado. Resolva esse
problema com um algoritmo cuja complexidade é O(n). */

package main

import "fmt"

const M = 10

func init() {
	fmt.Println("------------------------------")
	fmt.Println("Estrutura de Dados - AC5")
	fmt.Println("------------------------------")
	fmt.Println("Algoritmo c/ complexidade O(n)")
	fmt.Println("------------------------------")
}

func main() {
	arrayOrdenado := [M]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	r1, r2 := encontraPar(arrayOrdenado, 15)
	r3, r4 := encontraPar(arrayOrdenado, 20)
	fmt.Printf("(%d, %d)\n", r1, r2)
	fmt.Printf("(%d, %d)\n", r3, r4)

}

func encontraPar(arr [M]int, valorAlvo int) (int, int) {
	var valorInferior, valorSuperior int = 0, M - 1

	for valorInferior < valorSuperior {
		if (arr[valorInferior] + arr[valorSuperior]) == valorAlvo {
			return arr[valorInferior], arr[valorSuperior]
		}
		if (arr[valorInferior] + arr[valorSuperior]) > valorAlvo {
			valorSuperior = valorSuperior - 1
		} else {
			valorInferior = valorInferior + 1
		}

	}
	return -1, -1
}