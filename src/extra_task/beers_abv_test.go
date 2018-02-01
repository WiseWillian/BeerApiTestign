package main

import (
	"testing"
	"strconv"
	"io/ioutil"
	"net/http"
)

const endpoint = "http://api.brewerydb.com/v2/beers/?key=" //O endpoint da Api a ser consultado
const apiKey = "47705820af1e5f9f31c6700101bc6494" //A chave da api cadastrada pelo desenvolvedor

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

func makeRequest(style int, page int) ([]byte, error) {
	var request string = endpoint + apiKey + "&styleId=" + strconv.Itoa(style)
	
	if page != 0 {
		request += "&p=" + strconv.Itoa(page)
	}

	resp, err_req := http.Get(request) //Faz a requisição para o endpoint

	if err_req == nil {
		//Caso a requisição seja feita corretamente devemos processar a informacao
		body, err_read := ioutil.ReadAll(resp.Body) //Transformamos a resposta em []Byte
		return body, err_read
	}
	
	return nil, err_req
}

func TestAbv(t *testing.T){
	var style int = 1

	body, err := makeRequest(style, 0)

	if err != nil {
		t.Fatal(err)
	}
}