package main

import (
	"Banco_em_GO/contas"
	"fmt"
)

func PagarBoleto(conta verficarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

// Para de criar interfaces é preciso passar funções
type verficarConta interface {
	Sacar(valor float64) string
}

func main() {

	//Criando um cliente
	//clienteBruno := clientes.Titular{"Bruno", "12312", "Dev"}

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	//Para funcionar é preciso passar o endereço de onde está armazenado na memória
	PagarBoleto(&contaDoDenis, 60)

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 300)

	fmt.Println(contaDoDenis.ObterSaldo())
	fmt.Println(contaDaLuisa.ObterSaldo())
}
