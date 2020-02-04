package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

type gerente struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func gerenteConstructor(nome, email, senha string) gerente {
	var gerente gerente
	gerente.Nome = nome
	gerente.Email = email
	gerente.Senha = senha

	return gerente
}

func deleteGerenteFromDB(id string) (gerente gerente, err error) {
	var con *sql.DB = createCon()
	gerenteAntigo, err := getGerenteFromDB(id)
	if err != nil {
		log.Fatal(err)
		erro := errors.New(err.Error())
		return gerente, erro
	}
	resultado, err := con.Query(`DELETE FROM gerente WHERE gerente.id= '` + id + `'`)
	if err != nil {
		log.Fatal(err)
		erro := errors.New(err.Error())
		return gerente, erro
	}
	defer resultado.Close()
	for resultado.Next() {
		if err != nil {
			log.Fatal(err)
			erro := errors.New(err.Error())
			return gerente, erro
		}
		return gerenteAntigo, nil
	}
	erro := errors.New(err.Error())
	return gerente, erro
}

func getGerenteFromDB(id string) (gerente gerente, err error) {
	var con *sql.DB = createCon()

	resultado, err := con.Query(`SELECT nome, email FROM gerente where gerente.id= '` + id + `'`)
	if err != nil {
		log.Fatal(err)
	}
	defer resultado.Close()
	for resultado.Next() {
		err := resultado.Scan(&gerente.Nome, &gerente.Email)
		if err != nil {
			log.Fatal(err)
			erro := errors.New(err.Error())
			return gerente, erro
		}
		return gerente, nil
	}
	erro := errors.New(err.Error())
	return gerente, erro
}

func getGerentesFromDB() (gerentesSlice []gerente) {
	var con *sql.DB = createCon()

	resultado, err := con.Query("select nome, email from gerente")
	if err != nil {
		log.Fatal(err)
	}
	defer resultado.Close()
	for resultado.Next() {
		var gerente gerente
		err := resultado.Scan(&gerente.Nome, &gerente.Email)
		if err != nil {
			log.Fatal(err)
		} else {
			gerentesSlice = append(gerentesSlice, gerente)
		}
	}
	return
}

func insertGerente(gerente gerente) (result bool) {
	var con *sql.DB = createCon()
	fmt.Print(gerente)
	newPass := encrypting(gerente.Senha)
	gerente.Senha = newPass

	resultado, err := con.Query(`INSERT INTO gerente (nome, email, senha) VALUES ('` + gerente.Nome + `', '` + gerente.Email + `', '` + gerente.Senha + `');`)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer resultado.Close()
	return true
}

func encrypting(senha string) string {
	sha1Instance := sha1.New()
	sha1Instance.Write([]byte(senha))
	passwordCript := sha1Instance.Sum(nil)[:20]
	stringPasswordCript := hex.EncodeToString(passwordCript)
	return stringPasswordCript
}
