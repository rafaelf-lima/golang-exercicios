package main 

import "fmt"

type Contato struct{
	Nome string
	Email string
}

func (c *Contato) alteraEmail(email string){
	c.Email = email
}

func adicionaContato(contato Contato, contatos [5]Contato) [5]Contato {
	for i, c := range contatos{
		if (c == Contato{}){
			contatos[i] = contato
			break
		}
	}
	return contatos
}

func excluiContato(contatos [5]Contato) [5]Contato{
	var ind int = 0
	for i, c := range contatos{
		if (c == Contato{}){
			ind = i - 1
			break
		}
	contatos[ind] =  Contato{}
	}	
	return contatos
}



func main(){
	var n string
	var e string
	var listaContato [5]Contato

	fmt.Println("Seja bem-vindo!")

	 for {
		fmt.Println("Insira seu nome: ")
		fmt.Scanln(&n)
		fmt.Println("Insira seu e-mail:")
		fmt.Scanln(&e)
		listaContato = adicionaContato(Contato{Nome: n, Email: e}, listaContato)
		fmt.Println(listaContato)
	}







	
	
	






	






}