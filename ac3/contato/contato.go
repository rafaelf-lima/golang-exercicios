package contato

type Contato struct {
	Nome string
	Email string
}

func (c *Contato) AlteraEmail(novoEmail string){
	c.Email = novoEmail
}




