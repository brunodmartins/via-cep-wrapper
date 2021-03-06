package viacep

import (
	"BrunoDM2943/via-cep-wrapper/internal/constants/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

func (dto *ViaCepResponse) ToAddress() *domain.Address {
	result := &domain.Address{}
	result.City = dto.Localidade
	result.District = dto.Bairro
	result.State = dto.Uf
	result.Address = dto.Logradouro
	result.ZipCode = dto.Cep
	return result
}

var Client = http.DefaultClient
var Host = "https://viacep.com.br"

//GetLocation query the `viacep` HTTP API for a given zipCode
func GetLocation(zipCode string) (*ViaCepResponse, error) {
	response, err := Client.Get(fmt.Sprintf("%s/ws/%s/json/", Host, zipCode))
	if err != nil {
		return nil, err
	}
	bytes, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("erro buscando CEP na via CEP")
	}
	result := &ViaCepResponse{}
	_ = json.Unmarshal(bytes, result)
	if result.Cep == "" {
		return nil, errors.New("CEP not found on Via CEP")
	}
	return result, nil
}
