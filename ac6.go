package main

import "fmt"

const M = 5

func init(){
	fmt.Println("---------------------------------------------")
	fmt.Println("AC6 - Inserção ordenada (Estruturas de Dados)")
	fmt.Println("---------------------------------------------")
  }
  
func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, valor int) {
	if *n == M {
		fmt.Printf("Tentando inserir %d\n", valor)
		fmt.Println("Overflow")
	} else {
		if busca(*v, *n, valor) == -1 {
			fmt.Printf("Tentando inserir %d\n", valor)
			fmt.Printf("%d não encontrado, inserindo na lista\n", valor)
			v[*n] = valor
			*n++
		} else {
			fmt.Printf("O elemento %d já existe!\n", valor)
		}
	}
}

func busca(v [M]int, n, x int) int {
	i := 0
	v[n] = x
	for v[i] != x {
		i++
	}
	if i != n {
		return i
	}
	return -1

}
