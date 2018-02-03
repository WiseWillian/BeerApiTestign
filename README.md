# BeerApiTesting

Este repositório contém dois testes automatizados que chamam a Api do BreweryDB (http://www.brewerydb.com/developer), ambos com o intuito de validar informações presentes e a situação da Api.

## O primeiro teste (main_task)

O primeiro código consiste no teste de tipo de variável no retorno da Api. Para isso, foi necessário criar estruturas que recebessem os dados da resposta e os guardassem em campos do tipo **interface{}**. A escolha pelos campos deste tipo foi feita para que não houvesse "Type Assertion" logo após a transformação, assim foi possível testar o tipo de variável recebida em posições futuras.

As funções **fieldIsString(interface[})** e **fieldIsNumber(interface{})** fazem o retornam se o valor contido na interface tem o tipo correto. Além disso optou-se por guardar os erros notificados em um slice de erros (**[]error**), para assim poderem ser todos notificados em sequência

## O segundo teste (extra_task)

O segundo código consiste em um teste de teor alcoólico da bebida de acordo com seu Abv (Alcohol by volume). Cada bebida pertence a um estilo diferente, e este estilo presume qual o teor alcoólico máximo da bebida que pertece ao estilo (AbvMax), assim como o teor alcoólico mínimo (AbvMin). Assim, quando a bebida possuir um teor alcoólico fora do especificado por seu estilo, existe uma incongruência dos dados retornado que deve ser verificada.

## Instalando repositório

Clone o repositório no local desejado

```
git clone https://github.com/WiseWillian/BeerApiTesting/
```

### Caso seu sistema não seja linux_amd64

Navegue até o diretório clonado e execute: 

```
export GOPATH=$HOME/Local de clonagem/BeerApiTester
```

```
export GOBIN=$HOME/Local de clonagem/BeerApiTester/bin
```

```
go install src/beers_models
```

## Executando os testes

Para executar os testes, basta navegar até o diretório desejado.

```
cd src/main_task
```
ou

```
cd src/extra_task
```
e fazer o comando de execução de arquivo de testes do go: 

```
go test
```

## Tecnologias utilizadas

* [BreweryDB](http://www.brewerydb.com/developers) - Api web para pesquisa dos dados
* [Go](https://golang.org/doc/) - Linguagem de programação para execução dos testes

## Autores

* **Rafael Willian** - *Trabalho inicial* - [WiseWillian](https://github.com/WiseWillian)
