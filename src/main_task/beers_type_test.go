package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	"net/http"
	"encoding/json"
	"strconv"
	"errors"
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

func testCategoryTypes(category Category) []error {
	var test_errors []error
	
	if !fieldIsNumber(category.Id) && category.Id != nil{
		err := errors.New("O id da categoria (Category.Id) possui tipo diferente de int")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(category.Name) && category.Name != nil {
		err := errors.New("O nome da categoria (Category.Name) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(category.CreateDate) && category.CreateDate != nil {
		err := errors.New("A data de criação da categoria (Category.CreateDate) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	return test_errors
}

func testAvailableTypes(available Available) []error {
	var test_errors []error
	
	if !fieldIsNumber(available.Id) && available.Id != nil {
		fmt.Println("Passou")
		err := errors.New("O id de available (Available.Id) possui tipo diferente de int")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(available.Name) && available.Name != nil {
		err := errors.New("O nome de available (Available.Name) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(available.Description) && available.Description != nil {
		err := errors.New("A descrição de available (Available.Description) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	return test_errors
}

func testStyleTypes(style Style) []error {
	var test_errors []error
	
	if !fieldIsNumber(style.Id) && style.Id != nil {
		err := errors.New("O id do estilo (Style.Id) tipo diferente de int")
		test_errors = append(test_errors, err)
	}

	if !fieldIsNumber(style.CategoryId) && style.CategoryId != nil {
		err := errors.New("O id da categoria do estilo (Style.CategoryId) tipo diferente de int")
		test_errors = append(test_errors, err)
	}

	arr := testCategoryTypes(style.StyleCategory)

	for i := range arr {
		test_errors = append(test_errors, arr[i])
	}

	if !fieldIsString(style.Name) && style.Name != nil {
		err := errors.New("O nome do estilo (Style.Name) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.ShortName) && style.ShortName != nil {
		err := errors.New("A abreviação do nome do estilo (Style.ShortName) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.Description) && style.Description != nil {
		err := errors.New("A descrição do estilo (Style.Description) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.IbuMin) && style.IbuMin != nil {
		err := errors.New("O IBU mínimo do estilo (Style.IbuMin) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.IbuMax) && style.IbuMax != nil {
		err := errors.New("O IBU máximo do estilo (Style.IbuMax) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.AbvMin) && style.AbvMin != nil {
		err := errors.New("O ABV mínimo do estilo (Style.AbvMin) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.AbvMax) && style.AbvMax != nil {
		err := errors.New("O ABV máximo do estilo (Style.AbvMax) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.SrmMin) && style.SrmMin != nil {
		err := errors.New("O SRM mínimo do estilo (Style.SrmMin) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.SrmMax) && style.SrmMax != nil {
		err := errors.New("O SRM máximo do estilo (Style.SrmMax) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.OgMin) && style.OgMin != nil {
		err := errors.New("O OG mínimo do estilo (Style.OgMin) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.FgMin) && style.FgMin != nil {
		err := errors.New("O FG mínimo do estilo (Style.FgMin) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.FgMax) && style.FgMax != nil {
		err := errors.New("O FG máximo do estilo (Style.FgMax) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.CreateDate) && style.CreateDate != nil {
		err := errors.New("A data de criação do estilo (Style.CreateDate) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	if !fieldIsString(style.UpdateDate) && style.UpdateDate != nil {
		err := errors.New("A data de modificação do estilo (Style.CreateDate) possui tipo diferente de string")
		test_errors = append(test_errors, err)
	}

	return test_errors
}

func fieldIsString(field interface{}) bool {
	_, ok := field.(string)
	return ok
}

func fieldIsNumber(field interface{}) bool {
	_, ok := field.(float64)
	return ok
}

func getAllBeers(body []byte) (*BeerApiResponse, error){
	var beers = new (BeerApiResponse)
	err := json.Unmarshal(body, &beers)

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