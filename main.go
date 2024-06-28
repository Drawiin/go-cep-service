package main

import (
	"fmt"
	"os"
	"time"

	"github.com/drawiin/go-cep-service/service"
)

const defaulCep = "01153000"

func main() {
	args := os.Args
	var cep string
	if args == nil || len(args) < 2 {
		fmt.Println("Usando cep defaul")
		cep = defaulCep
	} else {
		cepArg := args[1]
		if len(cepArg) != 8 {
			fmt.Println("CEP invalido deve ter 8 digitos")
			return
		}
		cep = cepArg
	}

	fmt.Printf("Buscando CEP %s......\n", cep)

	viaCepService := service.ViaCepService{}
	brasilAPIService := service.BrasilAPIService{}

	viaCepChannel := make(chan service.CepResponse)
	brasilApiChannel := make(chan service.CepResponse)

	go func() {
		viaCepResponse, err := viaCepService.GetCep(cep)
		if err != nil {
			return
		}
		viaCepChannel <- viaCepResponse
	}()

	go func() {
		brasilApiResponse, err := brasilAPIService.GetCep(cep)
		if err != nil {
			return
		}
		brasilApiChannel <- brasilApiResponse
	}()

	select {
	case cepData := <-viaCepChannel:
		logResponse(cepData, "ViaCep")
	case cepData := <-brasilApiChannel:
		logResponse(cepData, "BrasilAPI")
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}

func logResponse(cepData service.CepResponse, service string) {
	fmt.Printf("\nCep recebido de %s\n\nCEP %s\nEstado: %s\nCidade: %s\nBairro: %s\nRua: %s\n", service, cepData.Cep, cepData.State, cepData.City, cepData.Neighborhood, cepData.Street)
}
