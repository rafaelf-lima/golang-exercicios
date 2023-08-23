package main

import (
	"bufio"
	"fmt"
	"os"
	"ac3/contato"
	"ac3/operacoes"
)

func main() {
	var ctts [5]contato.Contato
	var nome, email, opcao string
	//var ind int

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos: ")
	for {
		fmt.Println("Digite 1 para adicionar contato, 2 para remover contato, 3 para alterar e-mail, 4 para listar e qualquer outra coisa para sair")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			c := contato.Contato{Nome: nome, Email: email}
			fmt.Println(c)
			operacoes.AdicionaContato(c, &ctts)
			fmt.Println(ctts)
		case "2":
			operacoes.ExcluiContato(&ctts)
			fmt.Println(ctts)
		case "3":
			fmt.Println("Informe o índice do e-mail: ")
			// operacoes.EditaEmail(&ind, contatos)
			fmt.Println(ctts)
		case "4":
			fmt.Println("Contatos abaixo: ") 
			for i, c := range ctts{
				fmt.Println(i, c)
			}
		default:
			return
		}


	}

}