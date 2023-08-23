package main

import (
	"ac3/contato"
	"ac3/operacoes"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var ctts [5]contato.Contato
	var nome, email, opcao string
	var ind int

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
			//fmt.Println(ctts)
		case "2":
			operacoes.ExcluiContato(&ctts)
		case "3":
			fmt.Println("Lista dos contatos abaixo: ")
			for i, c := range ctts {
				fmt.Printf("(%d) Nome: %sE-mail: %s\n", i, c.Nome, c.Email)
				fmt.Println(" ")
			}
			fmt.Print("Informe o índice do e-mail a ser alterado: ")
			fmt.Scanln(&ind)
			operacoes.EditaEmail(ind, &ctts)
		case "4":
			fmt.Println("Lista dos contatos abaixo: ")
			for i, c := range ctts {
				fmt.Printf("(%d) Nome: %sE-mail: %s\n", i, c.Nome, c.Email)
				fmt.Println(" ")
			}
		default:
			return
		}

	}

}
