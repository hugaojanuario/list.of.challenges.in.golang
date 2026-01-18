package d8

import "fmt"

type Pagamento interface {
	Pagar(valor float64) bool
}

type CartaoCredito struct {
	Limite float64
}

func (cc *CartaoCredito) Pagar(valor float64) bool {
	if valor <= cc.Limite {
		fmt.Println("Pagamento realizado com sucesso!")
		return true
	} else {
		fmt.Println("Erro no pagamento!")
		return false
	}
}

type Pix struct {
	Saldo float64
}

func (p *Pix) Pagar(valor float64) bool {
	if valor <= p.Saldo && valor > 0 {
		fmt.Println("Pix realizado com sucesso!")
		return true
	} else {
		fmt.Println("Erro no pix!")
		return false
	}
}
