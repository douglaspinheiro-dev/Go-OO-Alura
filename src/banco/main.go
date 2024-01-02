package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	contaDoBruno := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Bruno",
			CPF:       "123.123.123-12",
			Profissao: "Desenvolvedor",
		},
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}
	contaDoBruno.Depositar(100)

	fmt.Println(contaDoBruno.ObterSaldo())

}
