package viacep

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocation_OK(t *testing.T) {
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
	Client = server.Client()
	Host = server.URL
	result, err := GetLocation("01001000")
	assert.Equal(t, result.Cep, "01001-000")
	assert.Nil(t, err)
}

func TestGetLocation_EmptyJSON_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("{}"))
		rw.WriteHeader(http.StatusOK)
	}))
	// Close the server when test finishes
	defer server.Close()
	Client = server.Client()
	Host = server.URL
	_, err := GetLocation("01001-000")
	assert.NotNil(t, err)
}

func TestGetLocation_NOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	}))
	// Close the server when test finishes
	defer server.Close()
	Client = server.Client()
	Host = server.URL
	_, err := GetLocation("01001-000")
	assert.NotNil(t, err)
}

