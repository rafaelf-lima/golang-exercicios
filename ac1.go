/* Primeiros exercícios da AC de Estrutura de Dados

Aluno: Rafael Lima

*/


package main

import "fmt"

func main(){

	// Exercício 1 da AC - Basta trocar o número para ver se é primo
	resultado1 := ePrimo(7)
	fmt.Println(resultado1)

	// Exercício 2 da AC - Basta trocar o número limite para descobrir o "n-ésimo"
	resultado2 := fibo(4)
	fmt.Printf("\nO n-ésimo número é %v!", resultado2)

	// Exercício 3 da AC - Basta trocar o número do dia para sair o nome correspondente definido no GitHub da disciplina
	resultado3 := diaSemana(7)
	fmt.Printf("\nO dia da semana é %v!", resultado3)

	// Exercício 4 da AC - Basta trocar o ano para definir se é bissexto ou não
	resultado4 := eBissexto(1900)
	fmt.Printf("\nO ano é %v!\n", resultado4)
}


func ePrimo(num int) string{
	var contador int = 0
	for i := 2; i < num; i++{
		if num % i == 0 {
			contador++
			fmt.Print(i, " - ")
		}
	}

	if contador >= 1 {
		return "O número não é primo!"
	} else {
		return "O número é primo!"
	}
}

func fibo(ex2 int) int {
	var num1, num2 int = 1, 1
	var nesimo int = 0
	fmt.Print("Sequência: ")
	for i := 0; i < ex2; i++ {
		fmt.Print(num1, " ")
		nesimo = num1
		num1, num2 = num2, num1+num2
	} 
	return nesimo
}



func diaSemana(dia int) string{
	switch dia{
		case 1:
			return "domingo"
		case 2:
			return "segunda-feira"
		case 3:
			return "terça-feira"
		case 4:
			return "quarta-feira"		
		case 5:
			return "quinta-feira"
		case 6:
			return "sexta-feira"
		case 7:
			return "sábado"
		default:
			return "não foi encontrado"	
	}

}


func eBissexto(ano int) string{
	if ano % 4 == 0 {	
		if ano % 100 == 0 && ano % 400 != 0 {
			// fmt.Printf("\nO ano %v não é bissexto!\n", ano)
			return "não bissexto"
		} else {
			// fmt.Printf("\nO ano %v é bissexto!\n", ano)
			return "bissexto"
		}
	} else {
		// fmt.Printf("\nO ano %v não é bissexto!\n", ano)
		return "não bissexto"
	}
}