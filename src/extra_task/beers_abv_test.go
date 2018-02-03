package mytest

// Autor: Rafael Willian 
//
// Descrição: Esse código se trata de um teste automatizado que verifica se o
// teor alcoólico da bebida se adequa ao estilo na qual ela se encontra. Para isso
// comparamos o atributo 'abv'(Alcohol by volume) da bebida com os atributos 'abv_max'
// e 'abv_min' presentes no estilo da bebida.

import (
	"testing"
	"strconv"
	"io/ioutil"
	"net/http"
	"beers_models"
)

const endpoint = "http://api.brewerydb.com/v2/beers/?key=" //O endpoint da Api a ser consultado
const apiKey = "47705820af1e5f9f31c6700101bc6494" //A chave da api cadastrada pelo desenvolvedor

//Função que faz a requisição ao endpoint e retorna um []Byte
func makeRequest(style int, page int) ([]byte, error) {
	var request string = endpoint + apiKey + "&styleId=" + strconv.Itoa(style) + "&p=" + strconv.Itoa(page)

	resp, err_req := http.Get(request) //Faz a requisição para o endpoint

	if err_req == nil {
		//Caso a requisição seja feita corretamente devemos processar a informacao
		body, err_read := ioutil.ReadAll(resp.Body) //Transformamos a resposta em []Byte
		return body, err_read
	}
	
	return nil, err_req
}

//Função principal do teste, que verifica se o abv da bebida está dentro do estilo
func TestAbv(t *testing.T){
	var style int = 1 //O estilo a ser verificado

	body, err := makeRequest(style, 1)

	//Verificação de um possível erro na requisicao
	if err != nil {
		t.Fatal(err)
	}

	beers, err_parse := beers_models.GetAllBeers(body) //Transforma o array de bytes e objetos

	//Verificação de erro na transformação do body []Byte em beers BeerApiResponse
	if err_parse != nil {
		t.Error(err_parse)
	}

	for i := 1; i <= beers.NumberOfPages; i++ { //Loop que navega entre paginas da resposta
		for j := range beers.Data { //Loop que itera sobre todas as bebidas da resposta
			if beers.Data[j].Abv != "" { //Caso a bebida possua um indice ABV

				//Capta o ABV da bebida, ABV maximo e minimo do estilo em numero
				abv, _ := strconv.ParseFloat(beers.Data[j].Abv, 64) 
				max_abv, _ := strconv.ParseFloat(beers.Data[j].BeerStyle.AbvMax, 64)
				min_abv, _ := strconv.ParseFloat(beers.Data[j].BeerStyle.AbvMin, 64)

				//Verifica se existe uma incongruencia entre o teor alcoólico da bebida e do estilo
				if abv > max_abv {
					t.Error("Id da bebida: " + beers.Data[j].Id + "A quantidade de álcool da bebida (" + beers.Data[j].Abv + ") é maior do que a que seu estilo prevê: " + beers.Data[j].BeerStyle.AbvMax)
				} else if abv < min_abv {
					t.Error("Id da bebida: " + beers.Data[j].Id + "A quantidade de álcool da bebida (" + beers.Data[j].Abv + ") é menor do que a que seu estilo prevê: " + beers.Data[j].BeerStyle.AbvMin)
				}
			}
		}

		body, err := makeRequest(style, i + 1) //Faz a requisição da próxima página

		//Verifica se houve erro na requisição da proxima página
		if err != nil {
			t.Fatal(err)
		}

		//Transforma a proxima página em novo objeto BeerApiResponse
		beers, err_parse = beers_models.GetAllBeers(body) 
	}
}