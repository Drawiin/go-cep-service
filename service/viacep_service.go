package service

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const viaCepURL = "https://viacep.com.br/ws/{cep}/json"

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func UnmarshalViaCepResponse(data []byte) (ViaCepResponse, error) {
	var r ViaCepResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

type ViaCepService struct{}

func (s *ViaCepService) GetCep(cep string) (CepResponse, error) {
	url := strings.Replace(viaCepURL, "{cep}", cep, 1)
	res, err := http.Get(url)
	if err != nil {
		return CepResponse{}, err
	}
	defer res.Body.Close()
	
	json, err := io.ReadAll(res.Body)
	if err != nil {
		return CepResponse{}, err
	}

	ViaCepResponse, err := UnmarshalViaCepResponse(json)
	if err != nil {
		return CepResponse{}, err
	}

	return CepResponse{
		Cep:          ViaCepResponse.Cep,
		State:        ViaCepResponse.Uf,
		City:         ViaCepResponse.Localidade,
		Neighborhood: ViaCepResponse.Bairro,
		Street:       ViaCepResponse.Logradouro,
	}, nil
}
