package main

import (
	"bufio"
	"fmt"
	"os"
	"ac3/operacoes"
	"ac3/contato"
)

func main() {
	var ctts [5]contato.Contato
	var nome, email, opcao string
	//var ind int

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Seja bem-vindo a lista de contatos!")
	for {
		fmt.Print("-------- Digite --------\n(1) para adicionar contato\n(2) para remover contato\n(3) para alterar e-mail\n(4) para listar\nQualquer outra coisa para sair \n")
		fmt.Println("------------------------")
		fmt.Print("Entrada: ")
		fmt.Scanln(&opcao)
		fmt.Println("------------------------")

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			operacoes.AdicionaContato(contato.Contato{Nome: nome, Email: email}, &ctts)
			fmt.Println(ctts)
		case "2":
			operacoes.ExcluiContato(&ctts)
			fmt.Println(ctts)
		case "3":
			fmt.Println("Aqui está a lista dos contatos")
			for i, c := range ctts{
				fmt.Println(i, c)
			}
			fmt.Println("Informe o índice do e-mail: ")
		case "4":
			fmt.Println("Contatos abaixo: ") 
			for i, c := range ctts{
				fmt.Printf("(%d) Nome: %sEmail: %s\n", i, c.Nome, c.Email)
				fmt.Println(" ")
			}
		default:
			return
		}


	}

}