package service

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const brasilApiURL = "https://brasilapi.com.br/api/cep/v1/{cep}"

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func UnmarshalBrasilAPIResponse(data []byte) (BrasilAPIResponse, error) {
	var r BrasilAPIResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

type BrasilAPIService struct{}

func (s *BrasilAPIService) GetCep(cep string) (CepResponse, error) {
	url := strings.Replace(brasilApiURL, "{cep}", cep, 1)
	res, err := http.Get(url)
	if err != nil {
		return CepResponse{}, err
	}
	defer res.Body.Close()
	
	json, err := io.ReadAll(res.Body)
	if err != nil {
		return CepResponse{}, err
	}

	brasilAPIResponse, err := UnmarshalBrasilAPIResponse(json)
	if err != nil {
		return CepResponse{}, err
	}

	return CepResponse{
		Cep:          brasilAPIResponse.Cep,
		State:        brasilAPIResponse.State,
		City:         brasilAPIResponse.City,
		Neighborhood: brasilAPIResponse.Neighborhood,
		Street:       brasilAPIResponse.Street,
	}, nil
}
