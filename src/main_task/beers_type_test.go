package main

import (
	"testing"
	"net/http"
)

const endpoint = "http://api.brewerydb.com/v2/beers/" //O endpoint da Api a ser consultado
const apiKey = "key=47705820af1e5f9f31c6700101bc6494" //A chave da api cadastrada pelo desenvolvedor
const chave = "styleId=1" //A pesquisa sera baseada pelo estilo da bebida. O Id 1 se refere à 'Classic English-Style Pale Ale'

func TestMain(t *testing.T){
	resp, err := http.Get(endpoint + "?" + apiKey + "&" + chave) //Faz a requisição para o endpoint

	if err != nil { //Testa um eventual erro na requisição
		t.Error("Erro ao solicitar informacao via GET para o endpoint")
	}

	defer resp.Body.Close()
}