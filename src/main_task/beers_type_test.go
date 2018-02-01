package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	"net/http"
)

const endpoint = "http://api.brewerydb.com/v2/beers/?key=" //O endpoint da Api a ser consultado
const apiKey = "47705820af1e5f9f31c6700101bc6494" //A chave da api cadastrada pelo desenvolvedor
const chave = "styleId=1" //A pesquisa sera baseada pelo estilo da bebida. O Id 1 se refere à 'Classic English-Style Pale Ale'

func TestMain(t *testing.T) {
	resp, err_req := http.Get(endpoint + apiKey + "&" + chave) //Faz a requisição para o endpoint

	//Testa um eventual erro na requisição
	if err_req != nil { 
		fmt.Println("Erro ao contatar o Endpoint: " + endpoint)
		t.Error(err_req)
	} else {
		//Caso a requisição seja feita corretamente devemos processar a informacao
		body, err_read := ioutil.ReadAll(resp.Body) //Transformamos a resposta em []Byte

		//Testa um erro na leitura e transformação da resposta
		if err_read != nil {
			fmt.Println("Erro ao ler a resposta")
			t.Error(err_read)
		} else {
			
		}
	}

	defer resp.Body.Close()
}