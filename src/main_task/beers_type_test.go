package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	"net/http"
	"encoding/json"
)

const endpoint = "http://api.brewerydb.com/v2/beers/?key=" //O endpoint da Api a ser consultado
const apiKey = "47705820af1e5f9f31c6700101bc6494" //A chave da api cadastrada pelo desenvolvedor
const chave = "styleId=1" //A pesquisa sera baseada pelo estilo da bebida. O Id 1 se refere à 'Classic English-Style Pale Ale'

type Category struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CreateDate string `json:"createDate"`
}

type Available struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type Style struct {
	Id int `json:"id"`
	CategoryId int `json:"categoryId"`
	StyleCategory Category `json:"category"`
	Name string `json:"name"`
	ShortName string `json:"shortName"`
	Description string `json:"description"`
	IbuMin string `json:"ibuMin"`
	IbuMax string `json:"ibuMax"`
	AbvMin string `json:"abvMin"`
	AbvMax string `json:"abvMax"`
	SrmMin string `json:"srmMin"`
	SrmMax string `json:"srmMax"`
	OgMin string `json:"ogMin"`
	FgMin string `json:"fgMin"`
	FgMax string `json:"fgMax"`
	CreateDate string `json:"createDate"`
	UpdateDate string `json:"updateDate"`
}

type Beer struct {
	Id string `json:"id"`
	Name string `json:"name"`
	NameDisplay string `json:"nameDisplay"`
	Description string `json:"description"`
	Abv string `json:"abv"`
	Ibu string `json:"ibu"`
	AvailableId int `json:"availableId"`
	StyleId int `json:"styleId"`
	IsOrganic string `json:"isOrganic"`
	Status string `json:"status"`
	StatusDisplay string `json:"statusDisplay"`
	CreateDate string `json:"createDate"`
	UpdateDate string `json:"updateDate"`
	BeerAvailable Available `json:"available"`
	BeerStyle Style `json:"style"`
}

type BeerApiResponse struct {
	CurrentPage int `json:"currentPage"`
	NumberOfPages int `json:"numberOfPages"`
	TotalResults int `json:"totalResults"`
	Data []Beer `json:"data"`
}

func getAllBeers(body []byte) (*BeerApiResponse, error){
	var beers = new (BeerApiResponse)
	err := json.Unmarshal(body, &beers)

	/* Todas as structs predefinem o tipo de variavel que deve ser recebida (Name string por ex,),
	dessa forma, se o tipo da variavel recebida for diferente do tipo especificado, um erro do tipo 
	TypeMismatch será lançado pelo método json.Unmarshal, caso contrário, todas as informações foram 
	recebidas corretamente, conforme o especificado */

	return beers, err
}

func TestMain(t *testing.T) {
	resp, err_req := http.Get(endpoint + apiKey + "&" + chave) //Faz a requisição para o endpoint

	//Testa um eventual erro na requisição
	if err_req != nil { 
		fmt.Println("Erro ao contatar o Endpoint: " + endpoint)
		t.Fatal(err_req)
	}
	
	//Caso a requisição seja feita corretamente devemos processar a informacao
	body, err_read := ioutil.ReadAll(resp.Body) //Transformamos a resposta em []Byte

	//Testa um erro na leitura e transformação da resposta
	if err_read != nil {
		fmt.Println("Erro ao ler a resposta")
		t.Fatal(err_read)
	}

	beers, err_parse := getAllBeers(body) //Transforma o array de bytes e objetos

	if err_parse != nil {
		t.Error(err_parse)
	}

	for i := range beers.Data {
		t.Log(beers.Data[i].Name)
	}

	defer resp.Body.Close()
}