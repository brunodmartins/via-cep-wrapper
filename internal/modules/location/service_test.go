package location

import (
	"BrunoDM2943/via-cep-wrapper/internal/gateway/viacep"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const zipCode = "01001-000"

func TestSearchLocation_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		responseBody := `
		{
			"cep": "01001-000",
			"logradouro": "Praça da Sé",
			"complemento": "lado ímpar",
			"bairro": "Sé",
			"localidade": "São Paulo",
			"uf": "SP",
			"ibge": "3550308",
			"gia": "1004",
			"ddd": "11",
			"siafi": "7107"
		  }
		`
		rw.Write([]byte(responseBody))
		rw.WriteHeader(http.StatusOK)
	}))
	// Close the server when test finishes
	defer server.Close()
	viacep.Client = server.Client()
	viacep.Host = server.URL
	result, err := SearchLocation(zipCode)
	assert.Equal(t, result.ZipCode, zipCode)
	assert.Nil(t, err)
}

func TestSearchLocation_NOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	}))
	// Close the server when test finishes
	defer server.Close()
	viacep.Client = server.Client()
	viacep.Host = server.URL

	_, err := SearchLocation(zipCode)
	assert.NotNil(t, err)
}


