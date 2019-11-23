package main

import(
	"crypto/sha1"
	"fmt"
)
type Gerente struct {
	nome, descricao, senha string
}

func createGerente(gerente Gerente) result bool{
	fmt.Println(gerente.nome)
	sha1Instance := sha1.New()
	sha1Instance.Write([]byte(gerente.senha))
	passwordCript := sha1Instance.Sum(nil)
}
