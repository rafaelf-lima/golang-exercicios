package contato

type Contato struct {
	Nome string
	Email string
}


func (c *Contato) alteraEmail(novoEmail string){
	c.Email = novoEmail
}

