package main

import "fmt"

func main(){

	resultado1 := ehPrimo(7)
	fmt.Println(resultado1)

	resultado2 := fibo(5)
	fmt.Println(resultado2)

	resultado3 := diaSemana(9)
	fmt.Println(resultado3)
}


func ehPrimo(z int) string{
	var contador int = 0
	for i := 2; i < z; i++{
		if z % i == 0 {
			contador++
			fmt.Print(i, " - ")
		}
	}

	if contador >= 1 {
		return "Não primo"
	} else {
		return "Primo"
	}
}


func diaSemana(x int) string{
	switch x{
		case 1:
			return "Domingo"
		case 2:
			return "Segunda-feira"
		case 3:
			return "Terça-feira"
		case 4:
			return "Quarta-feira"		
		case 5:
			return "Quinta-feira"
		case 6:
			return "Sexta-feira"
		default:
			return "Erro"	
	}

}


func e_bissexto(ano int){
	

}