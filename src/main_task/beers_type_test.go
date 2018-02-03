package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	"net/http"
	"encoding/json"
	"strconv"
)

const endpoint = "http://api.brewerydb.com/v2/beers/?key=" //O endpoint da Api a ser consultado
const apiKey = "47705820af1e5f9f31c6700101bc6494" //A chave da api cadastrada pelo desenvolvedor

type Category struct {
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	CreateDate interface{} `json:"createDate"`
}

type Available struct {
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	Description interface{} `json:"description"`
}

type Style struct {
	Id interface{} `json:"id"`
	CategoryId interface{} `json:"categoryId"`
	StyleCategory Category `json:"category"`
	Name interface{} `json:"name"`
	ShortName interface{} `json:"shortName"`
	Description interface{} `json:"description"`
	IbuMin interface{} `json:"ibuMin"`
	IbuMax interface{} `json:"ibuMax"`
	AbvMin interface{} `json:"abvMin"`
	AbvMax interface{} `json:"abvMax"`
	SrmMin interface{} `json:"srmMin"`
	SrmMax interface{} `json:"srmMax"`
	OgMin interface{} `json:"ogMin"`
	FgMin interface{} `json:"fgMin"`
	FgMax interface{} `json:"fgMax"`
	CreateDate interface{} `json:"createDate"`
	UpdateDate interface{} `json:"updateDate"`
}

type Beer struct {
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	NameDisplay interface{} `json:"nameDisplay"`
	Description interface{} `json:"description"`
	Abv interface{} `json:"abv"`
	Ibu interface{} `json:"ibu"`
	AvailableId interface{} `json:"availableId"`
	StyleId interface{} `json:"styleId"`
	IsOrganic interface{} `json:"isOrganic"`
	Status interface{} `json:"status"`
	StatusDisplay interface{} `json:"statusDisplay"`
	CreateDate interface{} `json:"createDate"`
	UpdateDate interface{} `json:"updateDate"`
	BeerAvailable Available `json:"available"`
	BeerStyle Style `json:"style"`
}

type BeerApiResponse struct {
	CurrentPage interface{} `json:"currentPage"`
	NumberOfPages interface{} `json:"numberOfPages"`
	TotalResults interface{} `json:"totalResults"`
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
	var estilo int = 1
	resp, err_req := http.Get(endpoint + apiKey + "&styleId=" + strconv.Itoa(estilo)) //Faz a requisição para o endpoint

	//Testa um eventual erro na requisição
	if err_req != nil { 
		fmt.Println("Erro ao contatar o Endpoint: " + endpoint)
		t.Fatal(err_req)
	}
	
	fmt.Println(resp.TransferEncoding)

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